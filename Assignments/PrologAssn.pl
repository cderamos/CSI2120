%Q2
adj(a,b).
adj(a,g).
adj(b,c).
adj(b,i).
adj(c,d).
adj(d,e).
adj(d,j).
adj(e,l).
adj(f,g).
adj(g,h).
adj(h,i).
adj(i,j).
adj(j,k).
adj(k,l).

color(red).
color(yellow).
color(blue).

%2A
% We assign a color to each windowpane, regardless of adjacency.
colorset([], []) :- !.
colorset([H|T], [X|L]) :- adj(H,_), color(X), colorset(T, L).

%2B
% Adjacent windowpanes are of different colors. We check that all windowpane colours
% ,that are neighboring windowpanes, have different colours.
% We do this by using the adj() predicate.
% If 2 windowpanes ARE adjacent, we move on to the next windowpane
diffadjcolor(_,_, [], []) :- !.
diffadjcolor(W, C, [H|T], [H1|T1]) :- adj(W, H), C \= H1, diffadjcolor(W,C,T,T1).
diffadjcolor(W, C, [H|T], [_|T1]) :- \+adj(W, H), diffadjcolor(W,C,T,T1).

%2C
% This predicate produces all valid combinations for a window.
% A "valid combination" means that no 2 adjacent windowpanes have the same color
% We use our previous predicates, diffadjcolor() and colorset() for this predicate
valid([],[]) :- !.
valid([A|Gs], [B|Cs]) :- diffadjcolor(A, B, Gs,Cs), valid(Gs, Cs).
generate(Gs,Cs):-colorset(Gs,Cs),valid(Gs,Cs).

%Q3
%This will return the country with the most points, if there is a tie, it will display
%the countries with the same highest point
choice(marie, [peru,greece,vietnam]).
choice(jean, [greece,peru,vietnam]).
choice(sasha, [vietnam,peru,greece]).
choice(helena,[peru,vietnam,greece]).
choice(emma, [greece,peru,vietnam]).


%Q3
% Trip Calculationm given a list of names, each person has ranked the countries in terms of what the priority is to visit said
% country. Tally up the totals to see what decided vote will be (what country has the most points will be the one the they visit
calculate(_,[],_) :-!.
calculate(ListOfNames,[Country|ListOfCountries],[P1|Result]) :- calculateIndividual(ListOfNames,Country,P1),
                                                                calculate(ListOfNames,ListOfCountries,Result).
calculateIndividual([],_,0) :-!.
calculateIndividual([F|N],Country, Point ) :-choice(F, Countries),reverse(CountryReverse, Countries),
                                             nth1(Rank,CountryReverse, Country),
                                             calculateIndividual(N, Country, PP),
                                             Point is PP+Rank.

where([H|T], FinalDecision) :- choice(H, CountryChoices), calculate([H|T],CountryChoices,FinalTally),
                               max_list(FinalTally,MaxNum),!,nth1(X, FinalTally, MaxNum)
                               ,nth1(X, CountryChoices, FinalDecision).


%Q4 A
% Return true if a given number can be divided by any of the numbers in the list.
% Used mod for this predicate. If N mod H, H divides N and the mod is 0, then return true and end.
% If not, check the rest of the list.
divisible([], _) :- !, fail.
divisible([H|_], N) :- N mod H =:= 0, !.
divisible([_|T], N) :- divisible(T,N).
