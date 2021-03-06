// source: proto/messages.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.ClientEvent', null, global);
goog.exportSymbol('proto.ClientEvent.DataCase', null, global);
goog.exportSymbol('proto.ClientEvent.NameChange', null, global);
goog.exportSymbol('proto.ClientEvent.Type', null, global);
goog.exportSymbol('proto.Event', null, global);
goog.exportSymbol('proto.Event.Client', null, global);
goog.exportSymbol('proto.Event.Client.ConnectionStatus', null, global);
goog.exportSymbol('proto.Event.Client.State', null, global);
goog.exportSymbol('proto.Event.DataCase', null, global);
goog.exportSymbol('proto.Event.Room', null, global);
goog.exportSymbol('proto.Event.ServerWelcome', null, global);
goog.exportSymbol('proto.Event.Type', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Event = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.Event.oneofGroups_);
};
goog.inherits(proto.Event, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.Event.displayName = 'proto.Event';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Event.Client = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Event.Client, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.Event.Client.displayName = 'proto.Event.Client';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Event.Room = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Event.Room, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.Event.Room.displayName = 'proto.Event.Room';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Event.ServerWelcome = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Event.ServerWelcome, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.Event.ServerWelcome.displayName = 'proto.Event.ServerWelcome';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ClientEvent = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.ClientEvent.oneofGroups_);
};
goog.inherits(proto.ClientEvent, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ClientEvent.displayName = 'proto.ClientEvent';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ClientEvent.NameChange = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ClientEvent.NameChange, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ClientEvent.NameChange.displayName = 'proto.ClientEvent.NameChange';
}

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.Event.oneofGroups_ = [[2,3,4]];

/**
 * @enum {number}
 */
proto.Event.DataCase = {
  DATA_NOT_SET: 0,
  CLIENT: 2,
  ROOM: 3,
  WELCOME: 4
};

/**
 * @return {proto.Event.DataCase}
 */
proto.Event.prototype.getDataCase = function() {
  return /** @type {proto.Event.DataCase} */(jspb.Message.computeOneofCase(this, proto.Event.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.Event.prototype.toObject = function(opt_includeInstance) {
  return proto.Event.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.Event} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Event.toObject = function(includeInstance, msg) {
  var f, obj = {
    type: jspb.Message.getFieldWithDefault(msg, 1, 0),
    client: (f = msg.getClient()) && proto.Event.Client.toObject(includeInstance, f),
    room: (f = msg.getRoom()) && proto.Event.Room.toObject(includeInstance, f),
    welcome: (f = msg.getWelcome()) && proto.Event.ServerWelcome.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Event}
 */
proto.Event.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.Event;
  return proto.Event.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Event} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Event}
 */
proto.Event.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.Event.Type} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 2:
      var value = new proto.Event.Client;
      reader.readMessage(value,proto.Event.Client.deserializeBinaryFromReader);
      msg.setClient(value);
      break;
    case 3:
      var value = new proto.Event.Room;
      reader.readMessage(value,proto.Event.Room.deserializeBinaryFromReader);
      msg.setRoom(value);
      break;
    case 4:
      var value = new proto.Event.ServerWelcome;
      reader.readMessage(value,proto.Event.ServerWelcome.deserializeBinaryFromReader);
      msg.setWelcome(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Event.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.Event.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Event} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Event.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getClient();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.Event.Client.serializeBinaryToWriter
    );
  }
  f = message.getRoom();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.Event.Room.serializeBinaryToWriter
    );
  }
  f = message.getWelcome();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.Event.ServerWelcome.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.Event.Type = {
  CLIENT_JOINED: 0,
  CLIENT_LEFT: 1,
  CLIENT_UPDATED: 2,
  SERVER_WELCOME: 3
};




if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.Event.Client.prototype.toObject = function(opt_includeInstance) {
  return proto.Event.Client.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.Event.Client} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Event.Client.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    connectionstatus: jspb.Message.getFieldWithDefault(msg, 3, 0),
    state: jspb.Message.getFieldWithDefault(msg, 4, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Event.Client}
 */
proto.Event.Client.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.Event.Client;
  return proto.Event.Client.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Event.Client} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Event.Client}
 */
proto.Event.Client.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {!proto.Event.Client.ConnectionStatus} */ (reader.readEnum());
      msg.setConnectionstatus(value);
      break;
    case 4:
      var value = /** @type {!proto.Event.Client.State} */ (reader.readEnum());
      msg.setState(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Event.Client.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.Event.Client.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Event.Client} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Event.Client.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getConnectionstatus();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.Event.Client.ConnectionStatus = {
  OFFLINE: 0,
  ONLINE: 1
};

/**
 * @enum {number}
 */
proto.Event.Client.State = {
  JOINING: 0,
  JOINED: 1
};

/**
 * optional string id = 1;
 * @return {string}
 */
proto.Event.Client.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.Event.Client} returns this
 */
proto.Event.Client.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.Event.Client.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.Event.Client} returns this
 */
proto.Event.Client.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional ConnectionStatus connectionStatus = 3;
 * @return {!proto.Event.Client.ConnectionStatus}
 */
proto.Event.Client.prototype.getConnectionstatus = function() {
  return /** @type {!proto.Event.Client.ConnectionStatus} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.Event.Client.ConnectionStatus} value
 * @return {!proto.Event.Client} returns this
 */
proto.Event.Client.prototype.setConnectionstatus = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * optional State state = 4;
 * @return {!proto.Event.Client.State}
 */
proto.Event.Client.prototype.getState = function() {
  return /** @type {!proto.Event.Client.State} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.Event.Client.State} value
 * @return {!proto.Event.Client} returns this
 */
proto.Event.Client.prototype.setState = function(value) {
  return jspb.Message.setProto3EnumField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.Event.Room.prototype.toObject = function(opt_includeInstance) {
  return proto.Event.Room.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.Event.Room} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Event.Room.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    clientsMap: (f = msg.getClientsMap()) ? f.toObject(includeInstance, proto.Event.Client.toObject) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Event.Room}
 */
proto.Event.Room.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.Event.Room;
  return proto.Event.Room.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Event.Room} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Event.Room}
 */
proto.Event.Room.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = msg.getClientsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.Event.Client.deserializeBinaryFromReader, "", new proto.Event.Client());
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Event.Room.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.Event.Room.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Event.Room} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Event.Room.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getClientsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(2, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.Event.Client.serializeBinaryToWriter);
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.Event.Room.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.Event.Room} returns this
 */
proto.Event.Room.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * map<string, Client> clients = 2;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.Event.Client>}
 */
proto.Event.Room.prototype.getClientsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.Event.Client>} */ (
      jspb.Message.getMapField(this, 2, opt_noLazyCreate,
      proto.Event.Client));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.Event.Room} returns this
 */
proto.Event.Room.prototype.clearClientsMap = function() {
  this.getClientsMap().clear();
  return this;};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.Event.ServerWelcome.prototype.toObject = function(opt_includeInstance) {
  return proto.Event.ServerWelcome.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.Event.ServerWelcome} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Event.ServerWelcome.toObject = function(includeInstance, msg) {
  var f, obj = {
    clientid: jspb.Message.getFieldWithDefault(msg, 1, ""),
    room: (f = msg.getRoom()) && proto.Event.Room.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Event.ServerWelcome}
 */
proto.Event.ServerWelcome.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.Event.ServerWelcome;
  return proto.Event.ServerWelcome.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Event.ServerWelcome} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Event.ServerWelcome}
 */
proto.Event.ServerWelcome.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setClientid(value);
      break;
    case 2:
      var value = new proto.Event.Room;
      reader.readMessage(value,proto.Event.Room.deserializeBinaryFromReader);
      msg.setRoom(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Event.ServerWelcome.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.Event.ServerWelcome.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Event.ServerWelcome} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Event.ServerWelcome.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClientid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRoom();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.Event.Room.serializeBinaryToWriter
    );
  }
};


/**
 * optional string clientId = 1;
 * @return {string}
 */
proto.Event.ServerWelcome.prototype.getClientid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.Event.ServerWelcome} returns this
 */
proto.Event.ServerWelcome.prototype.setClientid = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional Room room = 2;
 * @return {?proto.Event.Room}
 */
proto.Event.ServerWelcome.prototype.getRoom = function() {
  return /** @type{?proto.Event.Room} */ (
    jspb.Message.getWrapperField(this, proto.Event.Room, 2));
};


/**
 * @param {?proto.Event.Room|undefined} value
 * @return {!proto.Event.ServerWelcome} returns this
*/
proto.Event.ServerWelcome.prototype.setRoom = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.Event.ServerWelcome} returns this
 */
proto.Event.ServerWelcome.prototype.clearRoom = function() {
  return this.setRoom(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.Event.ServerWelcome.prototype.hasRoom = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional Type type = 1;
 * @return {!proto.Event.Type}
 */
proto.Event.prototype.getType = function() {
  return /** @type {!proto.Event.Type} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.Event.Type} value
 * @return {!proto.Event} returns this
 */
proto.Event.prototype.setType = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional Client client = 2;
 * @return {?proto.Event.Client}
 */
proto.Event.prototype.getClient = function() {
  return /** @type{?proto.Event.Client} */ (
    jspb.Message.getWrapperField(this, proto.Event.Client, 2));
};


/**
 * @param {?proto.Event.Client|undefined} value
 * @return {!proto.Event} returns this
*/
proto.Event.prototype.setClient = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.Event.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.Event} returns this
 */
proto.Event.prototype.clearClient = function() {
  return this.setClient(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.Event.prototype.hasClient = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional Room room = 3;
 * @return {?proto.Event.Room}
 */
proto.Event.prototype.getRoom = function() {
  return /** @type{?proto.Event.Room} */ (
    jspb.Message.getWrapperField(this, proto.Event.Room, 3));
};


/**
 * @param {?proto.Event.Room|undefined} value
 * @return {!proto.Event} returns this
*/
proto.Event.prototype.setRoom = function(value) {
  return jspb.Message.setOneofWrapperField(this, 3, proto.Event.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.Event} returns this
 */
proto.Event.prototype.clearRoom = function() {
  return this.setRoom(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.Event.prototype.hasRoom = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional ServerWelcome welcome = 4;
 * @return {?proto.Event.ServerWelcome}
 */
proto.Event.prototype.getWelcome = function() {
  return /** @type{?proto.Event.ServerWelcome} */ (
    jspb.Message.getWrapperField(this, proto.Event.ServerWelcome, 4));
};


/**
 * @param {?proto.Event.ServerWelcome|undefined} value
 * @return {!proto.Event} returns this
*/
proto.Event.prototype.setWelcome = function(value) {
  return jspb.Message.setOneofWrapperField(this, 4, proto.Event.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.Event} returns this
 */
proto.Event.prototype.clearWelcome = function() {
  return this.setWelcome(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.Event.prototype.hasWelcome = function() {
  return jspb.Message.getField(this, 4) != null;
};



/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.ClientEvent.oneofGroups_ = [[2]];

/**
 * @enum {number}
 */
proto.ClientEvent.DataCase = {
  DATA_NOT_SET: 0,
  NAMECHANGE: 2
};

/**
 * @return {proto.ClientEvent.DataCase}
 */
proto.ClientEvent.prototype.getDataCase = function() {
  return /** @type {proto.ClientEvent.DataCase} */(jspb.Message.computeOneofCase(this, proto.ClientEvent.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ClientEvent.prototype.toObject = function(opt_includeInstance) {
  return proto.ClientEvent.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ClientEvent} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ClientEvent.toObject = function(includeInstance, msg) {
  var f, obj = {
    type: jspb.Message.getFieldWithDefault(msg, 1, 0),
    namechange: (f = msg.getNamechange()) && proto.ClientEvent.NameChange.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ClientEvent}
 */
proto.ClientEvent.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ClientEvent;
  return proto.ClientEvent.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ClientEvent} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ClientEvent}
 */
proto.ClientEvent.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.ClientEvent.Type} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 2:
      var value = new proto.ClientEvent.NameChange;
      reader.readMessage(value,proto.ClientEvent.NameChange.deserializeBinaryFromReader);
      msg.setNamechange(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ClientEvent.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ClientEvent.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ClientEvent} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ClientEvent.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getNamechange();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.ClientEvent.NameChange.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.ClientEvent.Type = {
  NAME_CHANGE: 0
};




if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ClientEvent.NameChange.prototype.toObject = function(opt_includeInstance) {
  return proto.ClientEvent.NameChange.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ClientEvent.NameChange} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ClientEvent.NameChange.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ClientEvent.NameChange}
 */
proto.ClientEvent.NameChange.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ClientEvent.NameChange;
  return proto.ClientEvent.NameChange.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ClientEvent.NameChange} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ClientEvent.NameChange}
 */
proto.ClientEvent.NameChange.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ClientEvent.NameChange.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ClientEvent.NameChange.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ClientEvent.NameChange} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ClientEvent.NameChange.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string Name = 1;
 * @return {string}
 */
proto.ClientEvent.NameChange.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.ClientEvent.NameChange} returns this
 */
proto.ClientEvent.NameChange.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional Type type = 1;
 * @return {!proto.ClientEvent.Type}
 */
proto.ClientEvent.prototype.getType = function() {
  return /** @type {!proto.ClientEvent.Type} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.ClientEvent.Type} value
 * @return {!proto.ClientEvent} returns this
 */
proto.ClientEvent.prototype.setType = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional NameChange nameChange = 2;
 * @return {?proto.ClientEvent.NameChange}
 */
proto.ClientEvent.prototype.getNamechange = function() {
  return /** @type{?proto.ClientEvent.NameChange} */ (
    jspb.Message.getWrapperField(this, proto.ClientEvent.NameChange, 2));
};


/**
 * @param {?proto.ClientEvent.NameChange|undefined} value
 * @return {!proto.ClientEvent} returns this
*/
proto.ClientEvent.prototype.setNamechange = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.ClientEvent.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.ClientEvent} returns this
 */
proto.ClientEvent.prototype.clearNamechange = function() {
  return this.setNamechange(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ClientEvent.prototype.hasNamechange = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto);
