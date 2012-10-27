package sample

import (
  "appengine"
  "appengine/datastore"
  "net/http"
  "time"
)

type Entry struct {
  Name string
  Time time.Time
}

type DataManager struct {
  context appengine.Context
}

func NewDataManager(request *http.Request) *DataManager {
  return &DataManager{
    context: appengine.NewContext(request),
  }
}

func (data *DataManager) store(name string) error {
  entry := Entry{name, time.Now()}
  key := datastore.NewIncompleteKey(data.context, "Entry", nil /* parent */)
  _, err := datastore.Put(data.context, key, &entry)
  return err
}

func (data *DataManager) load(name string) (times []time.Time, err error) {
  query := datastore.NewQuery("Entry").
    Filter("Name=", name).Order("-Time").Limit(20)
  var entries []Entry
  _, err = query.GetAll(data.context, &entries)

  if err != nil {
    return
  }

  for _, entry := range entries {
    times = append(times, entry.Time)
  }
  return
}
