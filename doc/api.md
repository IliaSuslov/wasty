## API /api/v1/

## Get waybill list [GET] /waybills
- number
- sort
- skip
- limit 

## Create waybill [POST] /waybills
```
{
    Number string
    Date date
    Desc string
}
```

return:
```
    waybill{
        ID primitive.ObjectID
        Number string
        Date time.Time
        Desc string
    }
```

## Get waybill by id [Get] /waybills/{id}

return:
```
    waybill{
        ID primitive.ObjectID
        Number string
        Date time.Time
        Desc string
        ...
    }
```
## update waybill by id [POST] /waybills/{id}
```
    waybill{
        Number *string
        Date *time.Time
        Desc *string
        *Car
        *Driver
        *Firm
        *AddJobs
        *RmJob int
        *AddExecutions
    }
```

## [GET] /drivers
Params:

- name
- sort
- skip
- limit 

## [GET] /cars

Params:

- name
- sort
- skip
- limit 

## [GET] /addresses

Params:

- name
- sort
- skip
- limit 