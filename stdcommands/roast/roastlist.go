package roast

import (
	"math/rand"
)

func randomRoast() string {
	return roasts[rand.Intn(len(roasts))]
}

//Roasts copied from roast-bot
var roasts = []string{
	"I'd like to roast you, but it looks like your genetics already did.",
	"You look like someone set your face on fire and then put it out with a hammer.",
	"The only thing attracted to you is gravity.",
	"You're not good looking enough to be a model, but you're not smart enough to be anything else.",
	"If you'd like to know what sexual position produces the ugliest babies, you should ask your mother.",
	"Can you speak a little louder? I can't hear you over the sound of how stupid you are.",
	"Why are you even talking to me? So your self-esteem can match your IQ?",
	"I'm not insulting you, I'm describing you.",
	"If you hide your big nose and shut your big mouth, people will think you are attractive and well-spoken.",
	"You're so ugly, when your mom dropped you off at school, she got a fine for littering.",
	"Some babies were dropped on their heads, but you were clearly thrown at a wall.",
	"They say opposites attract. If that's so, you will meet someone who is good-looking, intelligent, and cultured.",
	"I didn't hear you. I'm busy ignoring an annoying person.",
	"I don't know what your problem is, but I'll bet it's hard to pronounce.",
	"It must take a lot of flexibility to fit your foot in your mouth and your head up your ass at the same time.",
	"I don't have the time nor the crayons to explain things to you.",
	"I'd love to keep chatting with you, but I'd rather have a root canal.",
	"I bet you swim with a t-shirt on.",
	"You have all the charm and charisma of a burning orphanage.",
	"Your face is so oily that I'm surprised America hasn't invaded yet.",
	"If you were any dumber, someone would need to water you twice a week.",
	"If you were on fire and I had a cup of my own piss, I'd drink it.",
	"Do you still love nature, despite what it did to you?",
	"The thing I dislike most about your face is that I can see it.",
	"If B.S. was music, you'd be an orchestra.",
	"You look like a before picture.",
	"I've heard farts more intelligent than you.",
	"You have a perfect face for radio.",
	"They say that a million monkeys on a million typewriters will eventually produce the collected works of Shakespeare. If that theory is correct, I believe you will one day say something intelligent.",
	"You look like somebody stepped on a goldfish.",
	"I thought the trash got picked up last night. What are you still doing here?",
	"Looking the way you do must save a lot of money on Halloween.",
	"I'd love to continue talking with you, but my favorite commercial is on TV.",
	"I'd love to keep chatting with you, but right now I have to do literally anything else.",
	"Did you get a bowl of soup with that haircut?",
	"If you don't like what I say about you, it would be a good idea to improve yourself.",
	"Does being that ugly require a license?",
	"You could throw a rock at the ground and miss.",
	"There's no one in this world like you. Or at least I hope so.",
	"Even a duck is smarter than you.",
	"Did you cancel your barbecue? Because your grill is messed up.",
	"Some people make millions. You make memes.",
	"Did you forget to wipe, or is that your natural scent?",
	"I missed you this week, but my aim is improving.",
	"I'm surprised you've made it this far without being eaten.",
	"Your body looks like your head is inflating a water balloon.",
	"Your mother was a hamster.",
	"How do you make an idiot wait? I'll tell you later.",
	"If balls were dynamite, you wouldn't have enough to kill a fish.",
	"I'd like to roast you, but I'm too busy judging your choices.",
	"You are the worst part of everybody's day.",
	"If your face were scrambled it would improve your looks.",
	"I hope you don't feel the way you look.",
	"In the book of Who's Who, you are listed as What's That?",
	"It's surprising to me that a pig's bladder on a stick has gotten so far in life.",
	"Sorry. I'm on the toilet and I can only deal with one shit at a time.",
	"If you fell into a river it would be unfortunate, but if anyone pulled you out it would be a disaster.",
	"You are the discount version of whatever celebrity you look like.",
	"When you go to the dentist, the dentist needs anesthetic.",
	"The fact that you are still alive is evidence that natural disasters are poorly distributed.",
	"You are so dumb you can't fart and chew gum at the same time.",
	"I was going to give you a nasty look, but I see you already have one.",
	"Me think'st thou are a general offense and every man should beat thee.",
	"You’re the clearance bin version of a personality.",
	"It’s a parent’s job to raise their children right. So looking at you, it’s no wonder your dad quit after just one day.",
	"You’re so ugly, when you got robbed, the robbers made you wear their masks.",
	"You're so ugly your face makes onions cry!",
	"You're not completely useless; you can always serve as a bad example.",
	"They say people get what they deserve. In your case, it's a participation trophy.",
	"I'd offer you some gum, but your smile's got plenty of it.",
	"If you were spice, you'd be flour.",
	"You're a negative ten on a scale of 0 to 5.",
	"You look like a pipe cleaner with eyes.",
	"Somewhere out there is a tree, tirelessly producing oxygen so you can breathe. I think you owe it an apology.",
	"You’ve aged.",
	"You're an open book written for very dumb children.",
	"Don't talk out loud. You lower the IQ of the whole street.",
	"You must have been born on the highway, because that's where most accidents happen.",
	"If you were a trophy at the end of a race, I'd walk backwards.",
	"The jerk store called, and they're running out of you!",
	"Your garden is overgrown, and your cucumbers are soft.",
	"I have had wet shits better than your personality.",
	"You should photoshop your life with better decisions.",
	"You weren’t a miracle baby; you were a clerical error.",
	"Your parents looked to see if they could return you.",
	"Whatever doesn't kill you, disappoints me.",
	"Take my lowest priority and put yourself beneath it.",
	"I don't hate you, but if you were drowning, I would give you a high five.",
	"Unless your name is Google, stop acting like you know everything.",
	"I told my therapist about you; she didn't believe me.",
	"I could eat alphabet soup and shit out a smarter statement than whatever you just said.",
	"You're as useless as the 'ueue' in 'queue.'",
	"You’re the reason the gene pool needs a lifeguard.",
	"Talking to you is like stepping on a leaf in autumn and hearing no crunch—disappointment.",
	"Even a Roomba has more direction in life than you do.",
	"You’re the kind of person who could lose a game of tic-tac-toe against yourself.",
}
