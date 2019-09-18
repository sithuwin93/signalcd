import 'package:ui/src/api/api.dart';

class API {
  API() {
    ApiClient client = ApiClient(basePath: "/api/v1");

    this.agents = AgentsApi(client);
    this.deployments = DeploymentsApi(client);
    this.pipelines = PipelineApi(client);
  }

  AgentsApi agents;
  DeploymentsApi deployments;
  PipelineApi pipelines;
}
