exports.handler = function(event, context, callback){
    var response = event.response;
    const updated = { "response": response + "World!" };
    callback(null, updated);
};
