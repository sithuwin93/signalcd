import 'package:angular/angular.dart';
import 'package:ui/agents_component.dart';
import 'package:ui/deployments_component.dart';
import 'package:ui/pipelines_component.dart';

@Component(
  selector: 'my-app',
  templateUrl: 'app_component.html',
  directives: [
    AgentsComponent,
    DeploymentsComponent,
    PipelinesComponent,
  ],
)
class AppComponent {}
