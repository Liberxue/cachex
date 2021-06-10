package main

import (
	"fmt"
	"time"

	"github.com/Liberxue/cachex"
	"github.com/golang/snappy"
)

func main() {
	src := []byte(`[{"_id":"60c1ecc92645556be032e56c","index":0,"guid":"80a94afc-419f-4a90-a7e3-0ec47c72a39e","isActive":false,"balance":"$2,622.41","picture":"http://placehold.it/32x32","age":39,"eyeColor":"blue","name":"Lynn Snow","gender":"male","company":"APEX","email":"lynnsnow@apex.com","phone":"+1 (963) 557-2200","address":"126 Ellery Street, Wilsonia, Tennessee, 6265","about":"Reprehenderit quis duis excepteur incididunt cillum. Do proident labore magna veniam tempor cupidatat aliquip. Fugiat quis et officia esse ad laboris labore reprehenderit aute. Laborum culpa sit dolore veniam minim velit voluptate commodo id dolore dolore sunt. Aute occaecat enim nostrud ad do et occaecat sint exercitation velit reprehenderit excepteur. Fugiat elit deserunt aliquip Lorem Lorem ex nulla est laborum adipisicing.\r\n","registered":"2020-03-16T10:58:30 -08:00","latitude":18.688394,"longitude":152.442755,"tags":["aliqua","proident","esse","est","nisi","quis","enim"],"friends":[{"id":0,"name":"Allen Ramos"},{"id":1,"name":"Willa Hawkins"},{"id":2,"name":"Eula Vang"}],"greeting":"Hello, Lynn Snow! You have 1 unread messages.","favoriteFruit":"strawberry"},{"_id":"60c1ecc95b80a0b4cdae7921","index":1,"guid":"ecb91f15-bf7e-46c1-84dc-17d8158f96d9","isActive":false,"balance":"$1,868.87","picture":"http://placehold.it/32x32","age":29,"eyeColor":"blue","name":"Dotson Walton","gender":"male","company":"ACLIMA","email":"dotsonwalton@aclima.com","phone":"+1 (841) 427-3534","address":"806 Preston Court, Kilbourne, North Dakota, 9120","about":"Consequat dolor duis est sunt et ea cillum fugiat eiusmod dolor culpa veniam. Fugiat eiusmod officia aliquip quis aliquip labore eiusmod mollit non nisi laborum ex. Commodo exercitation adipisicing veniam ex Lorem exercitation commodo ad. Elit amet officia occaecat aute laborum. Sunt in dolore commodo velit ad consectetur in et. Adipisicing aliqua reprehenderit aliqua tempor quis voluptate velit.\r\n","registered":"2015-06-27T04:58:18 -08:00","latitude":77.896096,"longitude":-71.175848,"tags":["officia","tempor","mollit","esse","veniam","anim","eu"],"friends":[{"id":0,"name":"Deena Jacobson"},{"id":1,"name":"Maxine Rice"},{"id":2,"name":"Douglas Morgan"}],"greeting":"Hello, Dotson Walton! You have 6 unread messages.","favoriteFruit":"banana"},{"_id":"60c1ecc9bdea064e7e9572f1","index":2,"guid":"6da7dd16-f72f-44fd-9110-1c60230e3e06","isActive":false,"balance":"$1,058.65","picture":"http://placehold.it/32x32","age":39,"eyeColor":"brown","name":"Tamika Hanson","gender":"female","company":"ENVIRE","email":"tamikahanson@envire.com","phone":"+1 (876) 584-3943","address":"877 Pineapple Street, Sandston, New Mexico, 8731","about":"Reprehenderit esse aliqua sit tempor. Sit proident qui Lorem Lorem mollit amet laboris et laboris ea Lorem. Mollit irure ea amet aliqua duis id ea est velit consectetur consectetur excepteur velit. Adipisicing ut excepteur in dolor duis laboris cillum magna magna ad dolor quis nisi.\r\n","registered":"2014-12-30T09:49:35 -08:00","latitude":-85.562349,"longitude":65.620803,"tags":["eu","Lorem","eu","nisi","duis","ex","officia"],"friends":[{"id":0,"name":"Fletcher Dejesus"},{"id":1,"name":"Young Farrell"},{"id":2,"name":"Mariana Sawyer"}],"greeting":"Hello, Tamika Hanson! You have 1 unread messages.","favoriteFruit":"apple"},{"_id":"60c1ecc9771018819b2f11a4","index":3,"guid":"fa149d1a-1198-4f56-a4a2-e0e6db48ff2e","isActive":true,"balance":"$3,277.20","picture":"http://placehold.it/32x32","age":29,"eyeColor":"green","name":"Sondra Murray","gender":"female","company":"DENTREX","email":"sondramurray@dentrex.com","phone":"+1 (910) 454-3037","address":"904 Dorchester Road, Cassel, Massachusetts, 8071","about":"Nulla duis esse voluptate aute. Aliqua dolor occaecat velit amet amet laboris enim nostrud nostrud sunt nisi est aliquip. Commodo pariatur ex pariatur excepteur cillum dolor enim reprehenderit ipsum tempor. Tempor aliqua quis Lorem amet irure id consequat.\r\n","registered":"2014-07-17T02:39:26 -08:00","latitude":-11.12054,"longitude":9.591102,"tags":["incididunt","reprehenderit","do","proident","laborum","adipisicing","in"],"friends":[{"id":0,"name":"Arnold Blevins"},{"id":1,"name":"Nanette May"},{"id":2,"name":"Theresa Franks"}],"greeting":"Hello, Sondra Murray! You have 10 unread messages.","favoriteFruit":"apple"},{"_id":"60c1ecc9f3d97c4b201a279f","index":4,"guid":"d887adc3-66f4-4f0a-9f14-5268636f25b0","isActive":true,"balance":"$1,688.26","picture":"http://placehold.it/32x32","age":34,"eyeColor":"brown","name":"Parker Guy","gender":"male","company":"QUIZKA","email":"parkerguy@quizka.com","phone":"+1 (901) 455-3451","address":"114 Dinsmore Place, Jenkinsville, Pennsylvania, 9811","about":"Elit cillum ut id ex mollit anim enim dolor consequat magna elit laboris dolor. Enim tempor duis mollit non labore dolore elit excepteur. Ex deserunt eu et ea dolor et reprehenderit id. Laborum commodo sunt occaecat esse cupidatat ex. Est nostrud aliquip deserunt enim magna et quis culpa id ea. Dolor est amet proident et proident. Ad dolor mollit qui fugiat sit cupidatat labore.\r\n","registered":"2020-08-20T04:16:23 -08:00","latitude":78.786832,"longitude":-120.907665,"tags":["consectetur","labore","duis","nostrud","voluptate","enim","occaecat"],"friends":[{"id":0,"name":"Stephens Cochran"},{"id":1,"name":"Bridgett Snider"},{"id":2,"name":"Camille Meadows"}],"greeting":"Hello, Parker Guy! You have 1 unread messages.","favoriteFruit":"banana"}]`)
	encoded := snappy.Encode(nil, src)
	c := cachex.NewCache(102400)
	var ch chan int
	ticker := time.NewTicker(time.Microsecond * 500)
	go func() {
		for range ticker.C {
			key := fmt.Sprintf("hello_%d", time.Now().UnixNano())
			err := c.Set(key, encoded)
			if err != nil {
				fmt.Println(err)
			}
			_, err = c.Get(key)
			if err != nil {
				fmt.Println(err)
			}
			// decoded, err := snappy.Decode(nil, encoded)
			// if err != nil {
			// 	log.Fatal(err)
			// }
		}
		ch <- 1
	}()
	<-ch

}
