# a simple Go webapp for deploying docker to Cloud Foundry or Lattice

hosted as a docker image at jbayer/lattice-app

To push to [Lattice](https://github.com/cloudfoundry-incubator/lattice) using [ltc](https://github.com/cloudfoundry-incubator/lattice/ltc):

```bash
ltc create lattice-app jbayer/lattice-app
```

To push to [Cloud Foundry](https://github.com/cf-release) using [cf](https://github.com/cloudfoundry/cli):

```bash
cf docker-push docker-app jbayer/lattice-app
```


### Endpoints

`/`: a simple landing page displaying the index and uptime  
`/env`: displays environment variables  
`/exit`: instructs Lattice to exit with status code 1  

### To rebuild the dockerimage:

```bash
./build.sh
```

Assumes you have the go toolchain (with the ability to cross-compile to different platforms) and docker installed and pointing at your docker daemon.
