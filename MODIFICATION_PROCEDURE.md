# Modification procedure

Steps to modify this code.

  1. Create a branch
  2. modify *version* and *Chart.yaml* accordingly
  3. insert your changes/improvements/fixes
  4. commit
  5. test

If test are ok:

  6. create the new helm package
    a. helm package helm-chart/nginxreverseproxy
    b. helm repo index .
  7. commit
  8. merge on *master*
  9. tag master with the new version

