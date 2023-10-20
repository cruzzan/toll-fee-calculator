# Coding assignment for Layer 10

Here are some initial things that I wanted to address in the software:
* Supplying pricing through config. Since this is an aspect of the operation that is very likely to change regularly, it's a nice thing to have in a malleable format.
* Not separating free rate from other pricing.
* Utilising the time lib a bit more to avoid comparison in other formats. This will help make the config easier to maintain too, since the format will be uniform.
* Add some testing, to ensure that a few potential bugs were not being triggered, mainly around the grace-period thing.

I had to take a few different routes on how to slice when rates apply, but in the end I feel like the day matchers and rate config turned out pretty well.

It should be decently easy to maintain now, since most changes should be configurable.

Vehicles, Holidays, and the service of calculation requests, I left fairly unchanged. Since it wasn't completely obvious to me how they would need to evolve.
Holidays should probably be fetched in batches from somewhere, instead of hardcoded. I'm sure there is a lib for that, too.

I also included an example config file. But I am not sure that a static JSON file would be a very good vehicle for that config. It is open to input errors...

Running main `go run .`

## Late edits 2023-10-20
* Corrected the order of the toll rate configs, so that the catchall isn't the first one.
* Correcting how rate configs handles end. So that it goes to the max tick of the minute.
* Set up matcher to apply the rate start and end to a new date based on the incoming instant. This fixes TZ preservation, and stops manipulation of the incoming instant.
* Adds initial sorting of the data coming to the calculator.

# Assignment: Toll fee calculator

A calculator for vehicle toll fees.

## Background

Our city has decided to implement toll fees in order to reduce traffic
congestion during rush hours.

This is the current draft of requirements:

- Fees will differ between 9 SEK and 22 SEK, depending on the time of day.
- The maximum fee for one day is 60 SEK.
- Only the highest fee should be charged for multiple passages within a 60
  minute period.
- Some vehicle types are fee-free.
- Fee-free days are; Saturdays, Sundays, holidays and day before holidays and
  the whole month of July. See [Transportstyrelsen][] for details.

## Your assignment

The last city-developer quit recently, claiming that this solution is
production-ready. You are now the new developer for our city - congratulations!

Your job is to deliver the code and from now on, you are the responsible
go-to-person for this solution. This is a solution you will have to put your
name on.

## Instructions

1.  Choose one of language alternatives available.
2.  Modify and re-factor the code as you see fit.
3.  Deliver your solution by e-mail or another suitable way.

## Help, I don't know Go, C, C#, Python, Java or JavaScript?!

No worries! We accept submissions in other languages as well, why not try it in
[Rust][] or [Kotlin][]?

[transportstyrelsen]: https://transportstyrelsen.se/sv/vagtrafik/Trangselskatt/Trangselskatt-i-goteborg/Tider-och-belopp-i-Goteborg/ "Trängselskatt i Göteborg - Transportstyrelsen"
[rust]: https://www.rust-lang.org/
[kotlin]: https://kotlinlang.org
