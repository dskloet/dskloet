package sample

import (
  "appengine"
  "appengine/datastore"
  "time"
)

type Entry struct {
  Name string
  Time time.Time
}

type DataManager struct {
  context appengine.Context
}

func NewDataManager(context appengine.Context) *DataManager {
  return &DataManager{
    context: context,
  }
}

func (data *DataManager) store(name string) error {
  entry := Entry{name, time.Now()}
  key := datastore.NewIncompleteKey(data.context, "Entry", nil /* parent */)
  _, err := datastore.Put(data.context, key, &entry)
  return err
}

func (data *DataManager) load(name string) (entries []Entry, err error) {
  query := datastore.NewQuery("Entry").
    Filter("Name=", name).Order("-Time").Limit(20)
  _, err = query.GetAll(data.context, &entries)

  if err != nil {
    return
  }
  return
}
