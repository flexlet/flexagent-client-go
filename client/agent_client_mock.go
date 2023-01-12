package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"time"

	"github.com/flexlet/flexagent-client-go/client/agent"
	"github.com/flexlet/flexagent-client-go/client/job"
	"github.com/flexlet/flexagent-client-go/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func NewTLSClient(host string, ca string, cert string, key string, insecure bool, formats strfmt.Registry) (*Agent, error) {
	tlsOpts := httptransport.TLSClientOptions{
		CA:                 ca,
		Key:                key,
		Certificate:        cert,
		InsecureSkipVerify: insecure,
	}
	client, err := httptransport.TLSClient(tlsOpts)
	if err != nil {
		return nil, err
	}
	transport := httptransport.NewWithClient(host, DefaultBasePath, DefaultSchemes, client)
	return New(transport, formats), nil
}

func NewTLSClientWithPEMBlock(host string, caPEMBlock []byte, certPEMBlock []byte, keyPEMBlock []byte, insecure bool, formats strfmt.Registry) (*Agent, error) {
	cfg := &tls.Config{}
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return nil, fmt.Errorf("tls client cert: %v", err)
	}
	cfg.Certificates = []tls.Certificate{cert}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caPEMBlock)
	cfg.RootCAs = caCertPool

	cfg.InsecureSkipVerify = insecure

	client := &http.Client{Transport: &http.Transport{TLSClientConfig: cfg}}
	transport := httptransport.NewWithClient(host, DefaultBasePath, DefaultSchemes, client)

	return New(transport, formats), nil
}

func (a *Agent) GetReadyStatus() (*models.ReadyStatus, error) {
	params := agent.NewReadyzParams()
	if resp, err := a.Agent.Readyz(params); err != nil {
		return nil, err
	} else {
		return resp.Payload, nil
	}
}

func (a *Agent) GetHealthStatus() (*models.HealthStatus, error) {
	params := agent.NewHealthzParams()
	if resp, err := a.Agent.Healthz(params); err != nil {
		return nil, err
	} else {
		return resp.Payload, nil
	}
}

func (a *Agent) SubmitJob(jobSpec *models.JobSpec, wait bool) (*models.Job, error) {
	params := job.NewSubmitParams()
	params.Spec = []*models.JobSpec{jobSpec}
	params.Wait = &wait
	if resp, err := a.Job.Submit(params); err != nil {
		return nil, err
	} else {
		return resp.Payload[0], nil
	}
}

func (a *Agent) WaitJob(jobUrn string, waitSec int) error {
	for i := 0; i < waitSec*100; i++ {
		if job, err := a.GetJob(jobUrn, nil, nil); err != nil {
			return err
		} else {
			// job complete
			if job.Status.State != models.JobStatusStateRunning && job.Status.State != models.JobStatusStateWaiting {
				return nil
			}
		}
		time.Sleep(time.Second * time.Duration(waitSec) / 100)
	}

	// job timeout
	return fmt.Errorf("job %s timeout", jobUrn)
}

func (a *Agent) GetJob(jobUrn string, outputLineLimit *int32, outputLineStart *int32) (*models.Job, error) {
	params := job.NewQueryParams()
	params.Urn = jobUrn
	params.OutputLineLimit = outputLineLimit
	params.OutputLineStart = outputLineStart
	if resp, err := a.Job.Query(params); err != nil {
		return nil, err
	} else {
		return resp.Payload, nil
	}
}

func (a *Agent) DeleteJob(jobUrn string) error {
	return nil
}
