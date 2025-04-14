package jsoncanvas

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
    reader := strings.NewReader(`{
	"nodes":[
		{"id":"a701d42981fd2a31","type":"group","x":-520,"y":-2120,"width":1420,"height":1189,"label":"Goals"},
		{"id":"debc3d22ab649578","type":"group","x":360,"y":720,"width":1092,"height":1460,"label":"Deployment"},
		{"id":"2c209d5866444c4b","type":"text","text":"#task-group\n# Configuration","x":420,"y":-391,"width":456,"height":121},
		{"id":"05ed6b9c8b27b7e1","type":"text","text":"#task\n# Create Production Dockerfile\n\nCreate a Dockerfile for building a production image.","x":957,"y":740,"width":475,"height":226},
		{"id":"42cf1665f2b0d5e4","type":"text","text":"#task-group\n# Deployment","x":380,"y":1035,"width":453,"height":110},
		{"id":"65af4bc8f8911308","type":"text","text":"#task\n# Define Docker Bake file\n\nSimplify build script by defining the Docker image in an HCL bake file.","x":957,"y":1000,"width":475,"height":230},
		{"id":"460b5f9c617ef3fb","type":"text","text":"#task\n# Define Kubernetes Manifests\n\nDefine Kubernetes Manifests required to deploy the application.","x":957,"y":1280,"width":475,"height":240},
		{"id":"591a79670de471cb","type":"text","text":"#task-group\n# Feature: Settings\n\nThe Settings Feature provides a way for defining GitHub Settings through a yaml configuration file. A developer should be able to set all non-destructive settings of a repository. The settings will be synchronised when a Commit is merged into the trunk branch.","x":437,"y":280,"width":423,"height":354},
		{"id":"9fb71593146bc5cc","type":"text","text":"# Zhāngyú māo\n\nZhāngyú māo (章鱼猫) or Octocat is a GitHub App that provides utilities for me to customise my GitHub User Experience.\n","x":-1080,"y":-210,"width":750,"height":250},
		{"id":"da7180e8ae0c4c20","type":"text","text":"# Work Items","x":-200,"y":80,"width":460,"height":80},
		{"id":"8dd61a11964bfb1f","type":"text","text":"#task-group\n# GitHub App Core\n\nGitHub Apps use two clients for interactions with a GitHub repository. (1) The App Client has limited functionality and is based on the App's ID, (2) the Installation Client which is based on the repository's installation ID and has access to all of the configured permissions; an App Client is required to instantiate the Installation Client.","x":437,"y":25,"width":423,"height":191},
		{"id":"eddce7db0217f264","type":"text","text":"#task\n# Define Build Workflow\n\nCreate a GitHub Workflow for building the Production Docker image and pushing it to the Container Registry.","x":957,"y":1580,"width":475,"height":240},
		{"id":"51b239e49ec486e8","type":"text","text":"#task\n# Create the GitHub App\n\nCreate the GitHub App and investigate whether creation and configuration of the app can be automated.","x":957,"y":1860,"width":475,"height":240},
		{"id":"7bd98ce7391182a1","type":"text","text":"#task\n# App Configuration\n\n- app_id\n- webhook_secret\n- private_key_path","x":980,"y":-800,"width":364,"height":240},
		{"id":"3ca383d96b1247e9","type":"text","text":"#task\n# App Proxy\n\nSet up a GitHub test repository, install the App on the repository and set up an ngrok/smee proxy to receive webhook events.","x":980,"y":-527,"width":364,"height":272},
		{"id":"3ba7d092b5d268b9","type":"text","text":"#task\n# Development Environment\nCreate a reproducible development setup using Nix, direnv and devenv.","x":980,"y":-221,"width":364,"height":301},
		{"id":"ade9689990e0901c","type":"text","text":"#metadata\nowner: SamWolfs\nrepo: zhangyu-mao","x":-1080,"y":116,"width":750,"height":244},
		{"id":"dea04b55c46f4346","type":"file","file":"The Twelve-Factor App.md","x":-500,"y":-2100,"width":620,"height":720},
		{"id":"1957fda5b8760306","type":"text","text":"# Twelve-Factor Implementation\n\n**I. Codebase**\n[Zhangyu Mao on GitHub](https://github.com/SamWolfs/zhangyu-mao)\n**II. Dependencies**\ngo.mod + investigate [Nix Flakes for Go programs](https://xeiaso.net/blog/nix-flakes-go-programs/)\n**III. Config**\n[Viper](https://github.com/spf13/viper)\n**IV. Backing services**\nN/A - For now, GitHub is the only backing service.\n**V. Build, release, run**\nDecide between GitOps or splitting GitHub Workflows.\n**VI. Processes**\n[Ergo](https://github.com/ergo-services/ergo)\n**VII. Port binding**\nYes\n**VIII. Concurrency**\n[Ergo](https://github.com/ergo-services/ergo)\n**IX. Disposability**\n[Ergo](https://github.com/ergo-services/ergo)\n**X. Dev/prod parity**\nHelm/GitOps/Tilt\n**XI. Logs**\nOpenTelemetry\n**XII. Admin processes**\nN/A\n\n","x":240,"y":-2100,"width":640,"height":720},
		{"id":"ddd0fb16ec518804","type":"text","text":"# Goals\n\n- Practice Golang\n- Become familiar with [Ergo](https://docs.ergo.services/)\n- Gain lower-level insights into GitHub Apps\n- Adhere to the [Twelve-factor App](12factor.net) best practices\n- Use Nix for a reproducible development environment","x":-420,"y":-1271,"width":460,"height":320}
	],
	"edges":[
		{"id":"a41b3a98f97a0056","fromNode":"ddd0fb16ec518804","fromSide":"top","toNode":"dea04b55c46f4346","toSide":"bottom"},
		{"id":"07b870bb96caf684","fromNode":"dea04b55c46f4346","fromSide":"right","toNode":"1957fda5b8760306","toSide":"left"},
		{"id":"c23ec1932dfdf2eb","fromNode":"9fb71593146bc5cc","fromSide":"right","toNode":"da7180e8ae0c4c20","toSide":"left"},
		{"id":"bc3bcfc9b480b8aa","fromNode":"da7180e8ae0c4c20","fromSide":"right","toNode":"2c209d5866444c4b","toSide":"left"},
		{"id":"9a28c1a76fb2139b","fromNode":"da7180e8ae0c4c20","fromSide":"right","toNode":"42cf1665f2b0d5e4","toSide":"left"},
		{"id":"cad82f0403fa4bb2","fromNode":"42cf1665f2b0d5e4","fromSide":"right","toNode":"05ed6b9c8b27b7e1","toSide":"left"},
		{"id":"3bd4400e70251d49","fromNode":"42cf1665f2b0d5e4","fromSide":"right","toNode":"65af4bc8f8911308","toSide":"left"},
		{"id":"139481cfdb144656","fromNode":"42cf1665f2b0d5e4","fromSide":"right","toNode":"460b5f9c617ef3fb","toSide":"left"},
		{"id":"c4f4608bc2dadeb8","fromNode":"42cf1665f2b0d5e4","fromSide":"right","toNode":"eddce7db0217f264","toSide":"left"},
		{"id":"bf472e538116668b","fromNode":"da7180e8ae0c4c20","fromSide":"right","toNode":"591a79670de471cb","toSide":"left"},
		{"id":"95b3e499b8e7cc13","fromNode":"da7180e8ae0c4c20","fromSide":"right","toNode":"8dd61a11964bfb1f","toSide":"left"},
		{"id":"ea0ccf19b0add6a0","fromNode":"2c209d5866444c4b","fromSide":"right","toNode":"7bd98ce7391182a1","toSide":"left"},
		{"id":"a46251aca4422ffa","fromNode":"2c209d5866444c4b","fromSide":"right","toNode":"3ca383d96b1247e9","toSide":"left"},
		{"id":"409dc132b8436a53","fromNode":"42cf1665f2b0d5e4","fromSide":"right","toNode":"51b239e49ec486e8","toSide":"left"},
		{"id":"b4b21f62d9deefe1","fromNode":"9fb71593146bc5cc","fromSide":"right","toNode":"ddd0fb16ec518804","toSide":"bottom"},
		{"id":"c2af1c5093b8e35c","fromNode":"2c209d5866444c4b","fromSide":"right","toNode":"3ba7d092b5d268b9","toSide":"left"},
		{"id":"3d943210d1e34b49","fromNode":"9fb71593146bc5cc","fromSide":"bottom","toNode":"ade9689990e0901c","toSide":"top"}
	]
}`)
	c, err := Decode(reader)
	require.NoError(t, err)
	fmt.Printf("%v", c)
}
