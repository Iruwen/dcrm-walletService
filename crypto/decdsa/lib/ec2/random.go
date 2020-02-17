/*
 *  Copyright (C) 2018-2019  Fusion Foundation Ltd. All rights reserved.
 *  Copyright (C) 2018-2019  caihaijun@fusion.org
 *
 *  This library is free software; you can redistribute it and/or
 *  modify it under the Apache License, Version 2.0.
 *
 *  This library is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  
 *
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package ec2 

import (
	"github.com/fsn-dev/dcrm-walletService/internal/common/math/random"
	"math/big"
	"time"
	"fmt"
)

var (
    SafePrime = make(chan *big.Int, 1000)
    RndInt = make(chan *big.Int, 1000)
)

func GenRandomSafePrime(length int) {

    /////tmp:set with const
    p1,_ := new(big.Int).SetString("88625664418799586874356962004865959785547839031996107242548047290486167353953657331351331208080542249656140168512748332634774039948640891157319400439768559406310998144033510222424018888076547189057455349935825269255274481159303548613626487902353694616182054651273353621573099119508591176180732381024939615643",10)
    SafePrime <-p1
    p2,_ := new(big.Int).SetString("87401419825340663460040109207636024489819852597825750089526559313377179426462816641223279522544592534348787386742046317830680361269939603721366618483477826928712515880829301278067339830286962135154436825955128973888418591265545000582273871980507982885605801232422264465693430004993253812233020752655348734603",10)
    SafePrime <-p2
    p3,_ := new(big.Int).SetString("72985346439917910990324494239972989926700971940502756081316073218425878220241272502553158989966898173499444116633661799751204973742018354368968343051074891067296400537009037174831360998635920810457897331656507392955202284250515584200206025296057705746893262452315422728351791191504340291908500417110130889127",10)
    SafePrime <-p3
    p4,_ := new(big.Int).SetString("87589984803689222342641559519379565869389641656963194348150678386007194163967771500292112091988828972021757838432733642890566078479448001852403286632908526127518034460270125495601529328841342960356689151399208978274696946380259695914713004876872197881088311692295225682235964090978442548593525553753584860563",10)
    SafePrime <-p4
    ////

    for {
	if len(SafePrime) < 4 { /////TODO  tmp:1000-->4
	    rndInt := <-RndInt
	    p := random.GetSafeRandomPrimeInt2(length/2,rndInt)
	    if p != nil {
		fmt.Println("================GenRandomSafePrime,p=%v====================",p)
		SafePrime <-p
		time.Sleep(time.Duration(1000000)) //1000 000 000 == 1s
	    }
	}

	////TODO tmp:1000-->4
	if len(SafePrime) == 4 {
	    break
	}
	//////
    }
}

func GenRandomInt(length int) {

    /////tmp:set with const
    time.Sleep(time.Duration(3000000000)) //1000 000 000 == 1s
    /////

    for {
	if len(RndInt) < 1000 {
	    ////TODO tmp:1000-->4
	    if len(SafePrime) == 4 {
		break
	    }
	    //////
	    p := random.GetSafeRandomInt(length/2)
	    RndInt <-p
	    
	    time.Sleep(time.Duration(1000000)) //1000 000 000 == 1s
	}
    }
}

