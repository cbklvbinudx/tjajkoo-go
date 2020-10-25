package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	basetext := `[Events]
//Background and Video events
//Storyboard Layer 0 (Background)
//Storyboard Layer 1 (Fail)
//Storyboard Layer 2 (Pass)
//Storyboard Layer 3 (Foreground)
Sprite,Foreground,Centre,"SB\approachcircle.png",0,0
 M,0,0,,200,400
 F,0,0,,0,1
 S,0,0,,0.51
 F,0,351327,,1,0
Sprite,Foreground,Centre,"SB\taikobigcircle.png",0,0
 M,0,0,,200,400
 F,0,0,,0,1
 S,0,0,,0.48
 C,0,0,,211,211,211
 F,0,351327,,1,0
Sprite,Foreground,Centre,"SB\approachcircle.png",0,0
 M,0,0,,280,400
 F,0,0,,0,1
 S,0,0,,0.51
 F,0,351327,,1,0
Sprite,Foreground,Centre,"SB\taikobigcircle.png",0,0
 M,0,0,,280,400
 F,0,0,,0,1
 S,0,0,,0.48
 C,0,0,,211,211,211
 F,0,351327,,1,0
Sprite,Foreground,Centre,"SB\approachcircle.png",0,0
 M,0,0,,360,400
 F,0,0,,0,1
 S,0,0,,0.51
 F,0,351327,,1,0
Sprite,Foreground,Centre,"SB\taikobigcircle.png",0,0
 M,0,0,,360,400
 F,0,0,,0,1
 S,0,0,,0.48
 C,0,0,,211,211,211
 F,0,351327,,1,0
Sprite,Foreground,Centre,"SB\approachcircle.png",0,0
 M,0,0,,440,400
 F,0,0,,0,1
 S,0,0,,0.51
 F,0,351327,,1,0
Sprite,Foreground,Centre,"SB\taikobigcircle.png",0,0
 M,0,0,,440,400
 F,0,0,,0,1
 S,0,0,,0.48
 C,0,0,,211,211,211
 F,0,351327,,1,0` + "\n"

	var osufilehalf []string // define the .osu file as a slice for easier detecting of quantity

	cwd, err := os.Getwd() // get current working directory
	if err != nil {
		fmt.Println("lol")
	}

	allfiles, err := ioutil.ReadDir(".") // list every file in directory
	if err != nil {
		fmt.Println("can't read directory")
	}

	for _, file := range allfiles { // get the name of the .osu file in directory
		if strings.HasSuffix(file.Name(), ".osu") {
			osufilehalf = append(osufilehalf, file.Name())
		}
	}

	if len(osufilehalf) > 1 {
		fmt.Println("too many osu files in one directory") // can't be more than 1 file in the directory
		panic(err)
	}

	osufile := cwd + `\` + osufilehalf[0] // add the current working directory path to the .osu file name, forming a proper path to the file

	fmt.Println(osufile)

	osuread, err := os.Open(osufile) // open the .osu file for reading
	if err != nil {
		fmt.Println("oops")
	}

	metadataline := bufio.NewScanner(osuread) // scanning for metadata that will allow to name the storyboard file

	var titlen string
	var artistn string
	var creatorn string

	for metadataline.Scan() {
		if strings.Contains(metadataline.Text(), "Title:") {
			titlen = metadataline.Text()[6:]
		} else if strings.Contains(metadataline.Text(), "Artist:") {
			artistn = metadataline.Text()[7:]
		} else if strings.Contains(metadataline.Text(), "Creator:") {
			creatorn = metadataline.Text()[8:]
		}
	}

	osbfile := artistn + " - " + titlen + " " + "(" + creatorn + ")" + ".osb" // concatenating the metadata, forming the name of the storyboard file

	fmt.Println(osbfile)

	osuread.Close()

	osbfilecreated, err := os.Create(osbfile) // create the storyboard file
	if err != nil {
		fmt.Println(err)
		return
	}

	appendosbopen, err := os.Open(osbfile) // open the storyboard file for appending
	if err != nil {
		fmt.Println("lol")
	}

	osbfilecreated.WriteString(basetext) // write the lines required for drawing the empty circles at the bottom

	osureadnew, err := os.Open(osufile) // open the .osu file for reading (can't use already created osuread variable)
	if err != nil {
		fmt.Println("oops")
	}

	oneline := bufio.NewScanner(osureadnew)

	loopcount := 0
	var elementcount int
	var hitobcount int

	for oneline.Scan() {
		if strings.Count(oneline.Text(), "[HitObjects]") == 1 {

			hitobcount++
		} else if strings.HasSuffix(oneline.Text(), ":") && hitobcount == 1 {

			elements := strings.Split(oneline.Text(), ",")

			colour, err := strconv.Atoi(elements[4])
			spinner := elements[5]
			timing, err := strconv.Atoi(elements[2])
			if err != nil {
				fmt.Println("wtf")
			}

			timingnotdone := timing - 1200
			timingdone := timing

			redright := `Sprite,Foreground,Centre,"SB\taikohitcircle.png",0,197` + "\n" +
				" MX,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",360\n" +
				" MY,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",-640,400\n" +
				" F,0," + strconv.Itoa(timingdone) + ",,1,0\n" +
				" S,0," + strconv.Itoa(timingdone) + ",,0.456\n" +
				" C,0," + strconv.Itoa(timingdone) + ",,235,69,44\n" +
				`Sprite,Foreground,Centre,"SB\taikohitcircleoverlay.png",0,197` + "\n" +
				" MX,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",360\n" +
				" MY,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",-640,400\n" +
				" F,0," + strconv.Itoa(timingdone) + ",,1,0\n" +
				" S,0," + strconv.Itoa(timingdone) + ",,0.456\n"

			redleft := `Sprite,Foreground,Centre,"SB\taikohitcircle.png",0,197` + "\n" +
				" MX,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",280\n" +
				" MY,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",-640,400\n" +
				" F,0," + strconv.Itoa(timingdone) + ",,1,0\n" +
				" S,0," + strconv.Itoa(timingdone) + ",,0.456\n" +
				" C,0," + strconv.Itoa(timingdone) + ",,235,69,44\n" +
				`Sprite,Foreground,Centre,"SB\taikohitcircleoverlay.png",0,197` + "\n" +
				" MX,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",280\n" +
				" MY,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",-640,400\n" +
				" F,0," + strconv.Itoa(timingdone) + ",,1,0\n" +
				" S,0," + strconv.Itoa(timingdone) + ",,0.456\n"

			blueright := `Sprite,Foreground,Centre,"SB\taikohitcircle.png",0,197` + "\n" +
				" MX,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",200\n" +
				" MY,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",-640,400\n" +
				" F,0," + strconv.Itoa(timingdone) + ",,1,0\n" +
				" S,0," + strconv.Itoa(timingdone) + ",,0.456\n" +
				" C,0," + strconv.Itoa(timingdone) + ",,67,142,172\n" +
				`Sprite,Foreground,Centre,"SB\taikohitcircleoverlay.png",0,197` + "\n" +
				" MX,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",200\n" +
				" MY,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",-640,400\n" +
				" F,0," + strconv.Itoa(timingdone) + ",,1,0\n" +
				" S,0," + strconv.Itoa(timingdone) + ",,0.456\n"

			blueleft := `Sprite,Foreground,Centre,"SB\taikohitcircle.png",0,197` + "\n" +
				" MX,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",440\n" +
				" MY,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",-640,400\n" +
				" F,0," + strconv.Itoa(timingdone) + ",,1,0\n" +
				" S,0," + strconv.Itoa(timingdone) + ",,0.456\n" +
				" C,0," + strconv.Itoa(timingdone) + ",,67,142,172\n" +
				`Sprite,Foreground,Centre,"SB\taikohitcircleoverlay.png",0,197` + "\n" +
				" MX,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",440\n" +
				" MY,0," + strconv.Itoa(timingnotdone) + "," + strconv.Itoa(timingdone) + ",-640,400\n" +
				" F,0," + strconv.Itoa(timingdone) + ",,1,0\n" +
				" S,0," + strconv.Itoa(timingdone) + ",,0.456\n"

			loopcount++
			elementcount = loopcount

			if spinner[0] != 48 { // if the hitobject is a spinner stay in place
				loopcount--
				elementcount++
			} else if colour == 12 || colour == 6 {
				osbfilecreated.WriteString(blueright + blueleft) // big blue hitobject
			} else if colour == 4 {
				osbfilecreated.WriteString(redright + redleft) // big red hitobject
			} else if loopcount%2 == 0 { // check if the hitobject should be drawn on the right side or the left side, making it alternate
				if colour == 8 || colour == 2 {
					osbfilecreated.WriteString(blueright)
				} else if colour == 0 {
					osbfilecreated.WriteString(redright)
				}
			} else if loopcount%2 != 0 {
				if colour == 8 || colour == 2 {
					osbfilecreated.WriteString(blueleft)
				} else if colour == 0 {
					osbfilecreated.WriteString(redleft)
				}
			}

			fmt.Println("Added " + strconv.Itoa(elementcount) + " elements.")
		} else {
			oneline.Scan() // if [HitObjects] not found keep scanning
		}
	}

	osbfilecreated.WriteString("//Storyboard Sound Samples") // this is required in order to work

	appendosbopen.Close()
	osuread.Close()
	osureadnew.Close()

	fmt.Println("\ngoinked")

}
