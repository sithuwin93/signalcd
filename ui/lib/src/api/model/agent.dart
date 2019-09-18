part of swagger.api;

class Agent {
  
  String name = null;
  

  bool ready = null;
  
  Agent();

  @override
  String toString() {
    return 'Agent[name=$name, ready=$ready, ]';
  }

  Agent.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    name =
        json['name']
    ;
    ready =
        json['ready']
    ;
  }

  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'ready': ready
     };
  }

  static List<Agent> listFromJson(List<dynamic> json) {
    return json == null ? new List<Agent>() : json.map((value) => new Agent.fromJson(value)).toList();
  }

  static Map<String, Agent> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, Agent>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new Agent.fromJson(value));
    }
    return map;
  }
}

