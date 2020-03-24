package value

import(
  "../entity"
)

type Group struct{
  Id int64
  RoomName string
  PeopleArray []entity.People
}
