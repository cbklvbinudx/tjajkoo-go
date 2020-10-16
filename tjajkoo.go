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

	var osufilehalf []string

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("lol")
	}

	allfiles, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("can't read directory")
	}

	for _, file := range allfiles {
		if strings.HasSuffix(file.Name(), ".osu") {
			osufilehalf = append(osufilehalf, file.Name())
		}
	}

	fmt.Println(osufilehalf)

	if len(osufilehalf) > 1 {
		fmt.Println("too many osu files in one directory")
		panic(err)
	}

	osufile := cwd + `\` + osufilehalf[0]

	fmt.Println(osufile)

	osuread, err := os.Open(osufile)
	if err != nil {
		fmt.Println("oops")
	}

	metadataline := bufio.NewScanner(osuread)

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

	osbfile := artistn + " - " + titlen + " " + "(" + creatorn + ")" + ".osb"

	fmt.Println(osbfile)

	osuread.Close()

	osbfilecreated, err := os.Create(osbfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	appendosbopen, err := os.Open(osbfile)
	if err != nil {
		fmt.Println("lol")
	}

	osbfilecreated.WriteString(basetext)

	osureadnew, err := os.Open(osufile)
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

			if spinner[0] != 48 {
				loopcount--
				elementcount++
			} else if colour == 12 || colour == 6 {
				osbfilecreated.WriteString(blueright + blueleft)
			} else if colour == 4 {
				osbfilecreated.WriteString(redright + redleft)
			} else if loopcount%2 == 0 {
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
			oneline.Scan()
		}
	}

	osbfilecreated.WriteString("//Storyboard Sound Samples")

	appendosbopen.Close()
	osuread.Close()

	fmt.Println("\ngoinked")

}
