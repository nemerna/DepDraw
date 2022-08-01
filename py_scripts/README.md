# POC-Scripts

the scripts existing in the directory were used for demonistrating and showing how we could parsea LucidChart Diagram into K8S manifests and applying those manifests.

in the other hand we have been demonistrating the other direction flow, which is cisualizing the deployed resources inside a new diagram using [Danielle Martinoli's project](https://github.com/dmartinol/openshift-topology-exporter).

---

## How to run

those scripts are expecting a CSV file path as an input, and this file should be representing a LucidChart exported diagram, additionally its expecting some output file name " use ./out.yaml" and it expects the number of manifests expected.

so using the [existing test-data](../test-data/test-diagram.csv), in order to convert it into K8S manifests the command should look like this:

`assuming the WorkingDir is py_scripts`

```
python3 main.py ../test-data/test-diagram.csv ./out.yaml 3

```