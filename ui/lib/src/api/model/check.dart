part of swagger.api;

class Check {
  
  String name = null;
  

  String image = null;
  

  List<String> imagePullSecrets = [];
  

  num duration = null;
  

  List<CheckEnvironment> environment = [];
  
  Check();

  @override
  String toString() {
    return 'Check[name=$name, image=$image, imagePullSecrets=$imagePullSecrets, duration=$duration, environment=$environment, ]';
  }

  Check.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    name =
        json['name']
    ;
    image =
        json['image']
    ;
    imagePullSecrets =
        (json['imagePullSecrets'] as List).map((item) => item as String).toList()
    ;
    duration =
        json['duration']
    ;
    environment =
      CheckEnvironment.listFromJson(json['environment'])
;
  }

  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'image': image,
      'imagePullSecrets': imagePullSecrets,
      'duration': duration,
      'environment': environment
     };
  }

  static List<Check> listFromJson(List<dynamic> json) {
    return json == null ? new List<Check>() : json.map((value) => new Check.fromJson(value)).toList();
  }

  static Map<String, Check> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, Check>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new Check.fromJson(value));
    }
    return map;
  }
}

