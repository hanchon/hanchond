package evmos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

type STRv1 struct {
	Denom    string
	Exponent int
	Alias    string
	Name     string
	Symbol   string
}

func (e *Evmos) CreateSTRv1Proposal(params STRv1) (string, error) {
	metadata := fmt.Sprintf(`
{
  "messages": [
    {
      "@type": "/cosmos.gov.v1.MsgExecLegacyContent",
      "authority": "evmos10d07y265gmmuvt4z0w9aw880jnsr700jcrztvm",
      "content": {
        "@type": "/evmos.erc20.v1.RegisterCoinProposal",
        "description": "IBC coin to erc-20",
        "metadata":[
        {
            "denom_units": [
              {
                "denom": "%s",
                "exponent": %d,
                "aliases": [
                  "%s"
                ]
              }
            ],
            "base": "%s",
            "display": "%s",
            "name": "%s",
            "symbol": "%s"
      }
      ]
      }
    }
  ],
  "deposit": "1000000000`+e.BaseDenom+`",
  "title": "STRv1 proposal",
  "summary": "Registering a new coin."
}`, params.Denom, params.Exponent, params.Alias, params.Denom, params.Alias, params.Name, params.Symbol)

	path := "/tmp/metadata.json"
	filesmanager.DoesFileExist(path)
	if err := filesmanager.SaveFile([]byte(metadata), path); err != nil {
		return "", fmt.Errorf("could not save the proposal to disk:%s", err.Error())
	}

	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"tx",
		"gov",
		"submit-proposal",
		path,
		"--keyring-backend",
		e.KeyringBackend,
		"--home",
		e.HomeDir,
		"--node",
		fmt.Sprintf("http://localhost:%d", e.Ports.P26657),
		"--from",
		e.ValKeyName,
		"--gas-prices",
		fmt.Sprintf("100%s", e.BaseDenom),
		"--gas-adjustment",
		"4",
		"--gas",
		"2000000",
		"-y",
	)

	out, err := command.CombinedOutput()
	return string(out), err
}

func (e *Evmos) VoteOnProposal(proposalID string, option string) (string, error) {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"tx",
		"gov",
		"vote",
		proposalID,
		option,
		"--keyring-backend",
		e.KeyringBackend,
		"--home",
		e.HomeDir,
		"--node",
		fmt.Sprintf("http://localhost:%d", e.Ports.P26657),
		"--from",
		e.ValKeyName,
		"--gas-prices",
		fmt.Sprintf("100%s", e.BaseDenom),
		"--gas-adjustment",
		"4",
		"-y",
	)

	out, err := command.CombinedOutput()
	if !strings.Contains(string(out), "code: 0") {
		return string(out), fmt.Errorf("transaction failed with code different than 0:%s", string(out))
	}
	hash := strings.Split(string(out), "txhash: ")
	if len(hash) > 1 {
		hash[1] = strings.TrimSpace(hash[1])
	}
	return hash[1], err
}

// VoteOnAllTheProposals returns a list of transactions hashes
func (e *Evmos) VoteOnAllTheProposals(option string) ([]string, error) {
	type ProposalsResponse struct {
		Proposals []struct {
			ProposalID string `json:"proposal_id"`
			Status     string `json:"status"`
		} `json:"proposals"`
	}
	// Query
	query := "cosmos/gov/v1beta1/proposals?pagination.limit=10&pagination.reverse=true"
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/%s", e.Ports.P1317, query))
	if err != nil {
		return []string{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return []string{}, fmt.Errorf("response not 200")
	}

	respbytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()
	var data ProposalsResponse
	if err := json.Unmarshal(respbytes, &data); err != nil {
		return []string{}, err
	}
	res := []string{}
	for _, v := range data.Proposals {
		if v.Status == "PROPOSAL_STATUS_VOTING_PERIOD" {
			// Vote
			out, err := e.VoteOnProposal(v.ProposalID, option)
			if err != nil {
				return []string{}, err
			}
			res = append(res, out)
		}
	}
	return res, nil
}

func (e *Evmos) CreateUpgradeProposal(upgradeHeight string) error {
	command := exec.Command( //nolint:gosec
		filesmanager.GetEvmosdPath(e.Version),
		"tx",
		"gov",
		"submit-legacy-proposal",
		"software-upgrade",
		"v19.0.0",
		"--title",
		"proposal",
		"--description",
		"description",
		"--upgrade-height",
		upgradeHeight,
		"--no-validate",
		"--deposit",
		"1000000000"+e.BaseDenom,
		"--keyring-backend",
		e.KeyringBackend,
		"--home",
		e.HomeDir,
		"--node",
		fmt.Sprintf("http://localhost:%d", e.Ports.P26657),
		"--from",
		e.ValKeyName,
		"--gas-prices",
		fmt.Sprintf("100%s", e.BaseDenom),
		"--gas-adjustment",
		"4",
		"--gas",
		"2000000",
		"-y",
	)

	out, err := command.CombinedOutput()
	fmt.Println(string(out))
	return err
}
