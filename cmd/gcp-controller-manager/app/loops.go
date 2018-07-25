/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"context"
	"crypto/x509"
	"fmt"
	"sort"
	"strings"

	"cloud.google.com/go/compute/metadata"
	"golang.org/x/oauth2"
	compute "google.golang.org/api/compute/v1"
	gcfg "gopkg.in/gcfg.v1"
	warnings "gopkg.in/warnings.v0"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"k8s.io/kubernetes/pkg/cloudprovider/providers/gce"
	"k8s.io/kubernetes/pkg/controller/certificates"
)

type controllerContext struct {
	client                      clientset.Interface
	sharedInformers             informers.SharedInformerFactory
	recorder                    record.EventRecorder
	gcpCfg                      GCPConfig
	clusterSigningGKEKubeconfig string
	done                        <-chan struct{}
}
type GCPConfig struct {
	ClusterName           string
	ProjectID             string
	Zones                 []string
	TokenSource           oauth2.TokenSource
	TPMEndorsementCACache *caCache
}

func loadGCPConfig(s *GCPControllerManager) (GCPConfig, error) {
	var cfg GCPConfig

	// Load gce.conf.
	gceConfig := struct {
		Global struct {
			ProjectID string `gcfg:"project-id"`
			TokenURL  string `gcfg:"token-url"`
			TokenBody string `gcfg:"token-body"`
		}
	}{}
	// ReadFileInfo will return warnings for extra fields in gce.conf we don't
	// care about. Wrap with FatalOnly to discard those.
	if err := warnings.FatalOnly(gcfg.ReadFileInto(&gceConfig, s.GCEConfigPath)); err != nil {
		return cfg, err
	}
	cfg.ProjectID = gceConfig.Global.ProjectID
	// Get the token source for GCE API.
	cfg.TokenSource = gce.NewAltTokenSource(gceConfig.Global.TokenURL, gceConfig.Global.TokenBody)

	// Get cluster zone from metadata server.
	zone, err := metadata.Zone()
	if err != nil {
		return cfg, err
	}
	// Extract region name from zone.
	if len(zone) < 2 {
		return cfg, fmt.Errorf("invalid master zone: %q", zone)
	}
	region := zone[:len(zone)-2]
	// Load all zones in the same region.
	client := oauth2.NewClient(context.Background(), cfg.TokenSource)
	cs, err := compute.New(client)
	if err != nil {
		return cfg, fmt.Errorf("creating GCE API client: %v", err)
	}
	allZones, err := compute.NewZonesService(cs).List(cfg.ProjectID).Do()
	if err != nil {
		return cfg, err
	}
	for _, z := range allZones.Items {
		if strings.HasPrefix(z.Name, region) {
			cfg.Zones = append(cfg.Zones, z.Name)
		}
	}
	if len(cfg.Zones) == 0 {
		return cfg, fmt.Errorf("can't find zones for region %q", region)
	}
	// Put master's zone first.
	sort.Slice(cfg.Zones, func(i, j int) bool { return cfg.Zones[i] == zone })

	cfg.ClusterName, err = metadata.Get("instance/attributes/cluster-name")
	if err != nil {
		return cfg, err
	}

	cfg.TPMEndorsementCACache = &caCache{
		rootCertURL: rootCertURL,
		interPrefix: intermediateCAPrefix,
		certs:       make(map[string]*x509.Certificate),
		crls:        make(map[string]*cachedCRL),
	}

	return cfg, nil
}

// loops returns all the control loops that the GCPControllerManager can start.
// We append GCP to all of these to disambiguate them in API server and audit
// logs. These loops are intentionally started in a random order.
func loops() map[string]func(*controllerContext) error {
	return map[string]func(*controllerContext) error{
		"certificate-approver": func(ctx *controllerContext) error {
			approver := newGKEApprover(ctx.gcpCfg, ctx.client)
			approveController := certificates.NewCertificateController(
				ctx.client,
				ctx.sharedInformers.Certificates().V1beta1().CertificateSigningRequests(),
				approver.handle,
			)
			go approveController.Run(5, ctx.done)
			return nil
		},
		"certificate-signer": func(ctx *controllerContext) error {
			signer, err := newGKESigner(ctx.clusterSigningGKEKubeconfig, ctx.recorder, ctx.client)
			if err != nil {
				return err
			}
			signController := certificates.NewCertificateController(
				ctx.client,
				ctx.sharedInformers.Certificates().V1beta1().CertificateSigningRequests(),
				signer.handle,
			)

			go signController.Run(5, ctx.done)
			return nil
		},
		"node-annotator": func(ctx *controllerContext) error {
			nodeAnnotateController, err := newNodeAnnotator(
				ctx.client,
				ctx.sharedInformers.Core().V1().Nodes(),
				ctx.gcpCfg.TokenSource,
			)
			if err != nil {
				return err
			}
			go nodeAnnotateController.Run(5, ctx.done)
			return nil
		},
	}
}

func loopNames() []string {
	names := make([]string, 0)
	for name := range loops() {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}