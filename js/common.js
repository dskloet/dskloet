var Ajax = function() {
  this.params = [];
};

Ajax.prototype.addParam = function(key, value) {
  if (value !== undefined) {
    this.params.push(encodeURIComponent(key) + '=' + encodeURIComponent(value));
  } else {
    this.params.push(encodeURIComponent(key));
  }
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

  var paramString = this.params.join('&');
  http.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  //http.setRequestHeader("Content-length", paramString.length);
  //http.setRequestHeader("Connection", "close");

  http.onreadystatechange = function() {
    if(http.readyState == 4) {
      if (http.status == 200) {
        onSuccess && onSuccess(http.responseText);
      } else {
        onFailure && onFailure();
      }
    }
  };
  http.send(paramString); 
};
