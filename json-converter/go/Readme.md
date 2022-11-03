### This Go Program converts the values in a CSV File to a JSON string, Hashes the JSON string and creates a new CSV with the Hash Column

## How to Use
* Clone the Repository and Navigate to the Go Folder
* Download Go https://go.dev/dl (My Apologies, I'm working on a build that won't require you to do this, but until then enjoy)
* Run the statement "go run main.go -csv `{path to your csv file}`"
* The csv flag above should contain a path to your csv file
* If the flag is not provided the default is an nft.csv file within the same folder (i.e nft.csv inside the go folder)
* ![h2](https://user-images.githubusercontent.com/68669102/199636636-ed9abf16-314b-42c5-b324-a31cacf59a0f.PNG)
* A New CSV called NewNft.csv will be generated each time the code is run and a new column "SHA256" will be added

##### The Input I used to build this is shown below (I did not leave it in the repo because i do not yet know the privacy policies attached to the HNG docs)

![h1](https://user-images.githubusercontent.com/68669102/199635609-1f0f2c34-b030-4d7d-b8a7-300dfd10aa90.PNG)

## Note: This Code is hardcoded in a lot of areas as the structure of the csv file I'm expecting is not clear yet.
