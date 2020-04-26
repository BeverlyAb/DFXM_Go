package data

import (
)

type Data struct {
    Msg int
    TID int
    CountID  int 	//used to match msg after refires, used as BOTH sending on compute and receiving
}