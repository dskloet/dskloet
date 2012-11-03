function getValue(id) {
  var element = document.getElementById(id);
  return element && element.value;
}

function onSearch(e) {
  e && e.preventDefault();

  var ajax = new Ajax();
  ajax.send('/search/' + escape(getValue('nameSearch')),
      onSearchResults, onError);
}

function onSearchResults(responseJson) {
  var response = JSON.parse(responseJson);
  if (!response.Entries) {
    return;
  }
  var list = document.createElement('ol');
  for (var i = 0; i < response.Entries.length; i++) {
    var item = document.createElement('li');
    item.innerHTML = response.Entries[i].Time;
    list.appendChild(item);
  }
  var container = document.getElementById('response');
  container.appendChild(list);
}

function onFormSubmit(e) {
  e.preventDefault();

  var ajax = new Ajax();
  var request = {first: getValue('first'), last: getValue('last')};
  ajax.addParam('nameJson', JSON.stringify(request));
  ajax.send('/hello', onResponse, onError);
}

function onResponse(response) {
  var responseElement = document.getElementById('response');
  responseElement.innerHTML = response;
}

function onError() {
  alert('Something went wrong');
}

function init() {
  var form = document.getElementById('formId');
  form.addEventListener('submit', onFormSubmit);
  var searchForm = document.getElementById('searchForm');
  searchForm && searchForm.addEventListener('submit', onSearch);
}

window.addEventListener('load', init);
