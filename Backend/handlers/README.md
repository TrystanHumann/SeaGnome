# [Endpoints - WIP]
_These endpoints are subjected to change_

##  Events API
___Bold__ query parameters are required_

### Event API Types:

__Result: EventPostBody__
```JSON
{
	"name": string
}
```

__Result: EventGetResponse__
```JSON
[
	{
		"id": int,
		"name": string,
		"complete": boolean,
		"created_date": timestamp
	},
    {
    	...
    }
]
```

| Endpoint      | Method           | Query Parameters  | Headers | Body | Response | Description | Auth |
| ------------- |:-------------:| :-----:| :-----:| :----:| :-----:| :----:| :---:|
| `/events`      | `GET` | N/A | N/A | N/A | `EventGetResponse` | Returns all events `DESC` by created_date | N/A |
| `/events`      | `POST` | N/A | N/A | `EventPostBody` |  N/A  | Creates a new event | N/A |
| `/events` | `PUT`     |    __id (int)__  && __completed (boolean)__| N/A | N/A | N/A | Updates an event complete status by the event ID | N/A |
| `/events`| `DELETE` | __id (int)__ | N/A | N/A | N/A | Deletes an event by its id (this should cascade to all foreign keys associated with this ID) | N/A |



# [Upload - WIP]
_These endpoints are subjected to change_

## Upload API
___Bold__ query parameters are required_

| Endpoint | Method | Query Parameters | Headers | Body | Response | Description | Auth | 
| ------------- |:-------------:| :-----:| :-----:| :----:| :-----:| :----:| :---:|
| `/upload` | `POST` | N/A | N/A | __uploadFile (file)__ && __gameID (int)__  (multiform) | N/A | Allows for excel file uploading for a specific game | N/A |
