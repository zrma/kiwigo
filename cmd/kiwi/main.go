package main

import (
	"fmt"
	"log"

	kiwi "github.com/codingpot/kiwigo"
)

func main() {
	kb := kiwi.NewBuilder("./ModelGenerator", 1 /*=numThread*/, kiwi.KIWI_BUILD_INTEGRATE_ALLOMORPH /*=options*/)

	err := kiwi.KiwiError()
	if err != "" {
		log.Panicln("err", err)
	}

	kb.AddWord("코딩냄비", "NNP", 0)

	k := kb.Build()
	defer k.Close() // don't forget to Close()!

	results, _ := k.Analyze("안녕하세요 코딩냄비입니다. 부글부글.", 1 /*=topN*/, kiwi.KIWI_MATCH_ALL)
	fmt.Println(results)
}
