#!/bin/bash

function gophersay {
# base64 encoded then gzipped golang gopher ascii art
# from here: https://gist.github.com/belbomemo/b5e7dad10fa567a5fe8a
gopherart="H4sIAJ73h1cAA32OSxLCIBBE95yiy40VCoNHYOsZggyruMnGPebs9pCPSSztED49bxpw+SEDIOo0
z4j8nbAyqhRpBYuIiHPibBJq8nQnOduHjSfriHhaHv5cq6kFjTk5hSfXHqYDQo+C0DRaKVj3A3AF
25kTszbGsV6loYj1Lr8809TOVTIIx8a4Tczr43Q1QO57iMzWYWjZBevhyGDk983s2o4q03v+QpUx
b+lcNw2vAQAA
"

say[0]="There is no friend as loyal as a book."
say[1]="There is nothing to writing. All you do is sit down at a typewriter and bleed."
say[2]="Happiness in intelligent people is the rarest thing I know."
say[3]="I love sleep. My life has the tendency to fall apart when I'm awake, you know?"
say[4]="The best way to find out if you can trust somebody is to trust them."
say[5]="Always do sober what you said you'd do drunk. That will teach you to keep your mouth shut."
say[6]="There is nothing noble in being superior to your fellow man; true nobility is being superior to your former self."
say[7]="The world breaks everyone, and afterward, many are strong at the broken places."
say[8]="The first draft of anything is shit."
say[9]="The most painful thing is losing yourself in the process of loving someone too much, and forgetting that you are special too."
say[10]="All you have to do is write one true sentence. Write the truest sentence that you know."
say[11]="I drink to make other people more interesting."
say[12]="When people talk, listen completely. Most people never listen."
say[13]="All good books are alike in that they are truer than if they had really happened and after you are finished reading one you will feel that all that happened to you and afterwards it all belongs to you: the good and the bad, the ecstasy, the remorse and sorrow, the people and the places and how the weather was. If you can get so that you can give that to people, then you are a writer."
say[14]="Every day is a new day. It is better to be lucky. But I would rather be exact. Then when luck comes you are ready."
say[15]="All thinking men are atheists."
say[15]="Courage is grace under pressure."
say[16]="Never think that war, no matter how necessary, nor how justified, is not a crime."
say[17]="But man is not made for defeat, A man can be destroyed but not defeated."
say[18]="Every man's life ends the same way. It is only the details of how he lived and how he died that distinguish one man from another."
say[19]="The world breaks every one and afterward many are strong at the broken places."
say[20]="Never confuse movement with action."
say[21]="A cat has absolute emotional honesty: human beings, for one reason or another, may hide their feelings, but a cat does not."
say[22]="You can't get away from yourself by moving from one place to another."
say[23]="Forget your personal tragedy. We are all bitched from the start and you especially have to be hurt like hell before you can write seriously. But when you get the damned hurt, use it-don't cheat with it."
say[24]="I’m not brave any more darling. I’m all broken. They’ve broken me."

# random
rand=$[$RANDOM % 25]

# Echo random saying and art
echo ' ------------------------'
echo " ${say["$rand"]}"
echo "$gopherart" | base64 --decode | gzip -d

}

gophersay
