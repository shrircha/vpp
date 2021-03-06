// Code generated by 'create-test-data.sh' on Thu Oct  4 15:14:27 PDT 2018. DO NOT EDIT.

package testdata

func getRawK8sNodeTestData() []string {
	return []string{

		`{
			"addresses": [
				{
					"address": "10.20.0.2",
					"type": 3
				},
				{
					"address": "k8s-master",
					"type": 1
				}
			],
			"name": "k8s-master",
			"node_info": {
				"Architecture": "amd64",
				"KubeProxyVersion": "v1.10.5",
				"OperatingSystem": "linux",
				"boot_ID": "0fc55bef-1478-4508-8d6b-5434e4c2f381",
				"container_runtime_version": "docker://18.6.1",
				"kernel_version": "4.4.0-21-generic",
				"kubelet_version": "v1.10.5",
				"machine_ID": "91550c3d3d1bca06c11d4f64575584db",
				"os_image": "Ubuntu 16.04 LTS",
				"system_UUID": "0C70EDA6-D715-473D-A3EE-405AEF96541A"
			},
			"pod_CIDR": "10.0.0.0/24"
		}`,
		`{
			"addresses": [
				{
					"address": "10.20.0.10",
					"type": 3
				},
				{
					"address": "k8s-worker1",
					"type": 1
				}
			],
			"name": "k8s-worker1",
			"node_info": {
				"Architecture": "amd64",
				"KubeProxyVersion": "v1.10.5",
				"OperatingSystem": "linux",
				"boot_ID": "f308b1c8-aeb9-45a5-aea4-47bb07f3df99",
				"container_runtime_version": "docker://18.6.1",
				"kernel_version": "4.4.0-21-generic",
				"kubelet_version": "v1.10.5",
				"machine_ID": "91550c3d3d1bca06c11d4f64575584db",
				"os_image": "Ubuntu 16.04 LTS",
				"system_UUID": "4DB7FFFD-595E-4CB9-8BC9-D20462DF30E8"
			},
			"pod_CIDR": "10.0.1.0/24"
		}`,
		`{
			"addresses": [
				{
					"address": "10.20.0.11",
					"type": 3
				},
				{
					"address": "k8s-worker2",
					"type": 1
				}
			],
			"name": "k8s-worker2",
			"node_info": {
				"Architecture": "amd64",
				"KubeProxyVersion": "v1.10.5",
				"OperatingSystem": "linux",
				"boot_ID": "4703280b-a0f9-4dfb-92bb-d3147f183935",
				"container_runtime_version": "docker://18.6.1",
				"kernel_version": "4.4.0-21-generic",
				"kubelet_version": "v1.10.5",
				"machine_ID": "91550c3d3d1bca06c11d4f64575584db",
				"os_image": "Ubuntu 16.04 LTS",
				"system_UUID": "6AEF96E5-C48F-4A14-BBA2-45CF523E6F30"
			},
			"pod_CIDR": "10.0.2.0/24"
		}`,
	}
}
