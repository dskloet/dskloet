var Ajax = function() {
  this.formData = new FormData();
};

Ajax.prototype.addParam = function(key, value) {
  this.formData.append(key, value || '');
};

Ajax.prototype.addField = function(field) {
  var element = document.getElementById(field);
  this.addParam(element.name, element.value);
};

Ajax.prototype.send = function(url, onSuccess, onFailure) {
  var http;
  if (window.XMLHttpRequest) {
    // Mozilla/Safari
    http = new XMLHttpRequest();
  } else if (window.ActiveXObject) {
    // IE
    http = new ActiveXObject("Microsoft.XMLHTTP");
  } else {
    onFailure && onFailure();
  }
  http.open('POST', url, true);

  http.onreadystatechange = function() {
    if(http.readyState == 4) {
      if (http.status == 200) {
        onSuccess && onSuccess(http.responseText);
      } else {
        onFailure && onFailure();
      }
    }
  };
  http.send(this.formData); 
};
