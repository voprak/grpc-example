# disable share button.
disable_snapshots()
# disable telemetry
analytics_settings(enable=False)
# print context to
print("context is: {}\n".format(k8s_context()))

# Tilt Extension to sync files without restarting container when required
load('ext://restart_process', 'docker_build_with_restart')

# configure go compilation command to run on host machine
local_resource(
  'greeter-server-compile',
  'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o greeter-server/build/greeter-server greeter-server/rpci/server.go',
  deps=['greeter-server/rpci/server.go', 'greeter-server/pb/service.pb.go'],
)

# build docker image and performs smart syncs of go binary without rebuilding the image when possible
docker_build_with_restart(
  'greeter-server-image',
  './greeter-server',
  entrypoint=['/app/build/greeter-server'],
  dockerfile='greeter-server/deployments/Dockerfile',
  only=[
    './build'
  ],
  live_update=[
    sync('./greeter-server/build', '/app/build'),
  ],
)
# watch kubernetes deployment
k8s_yaml('greeter-server/deployments/kubernetes.yaml')
# create/configure greeter-server kubernetes deployment resource
k8s_resource('greeter-server', port_forwards=59000, resource_deps=['greeter-server-compile'])
