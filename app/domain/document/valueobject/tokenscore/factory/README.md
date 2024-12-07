Okapi BM 25　　

${\displaystyle {\text{score}}(D,Q)=\sum _{i=1}^{n}{\text{IDF}}(q_{i})\cdot {\frac {f(q_{i},D)\cdot (k_{1}+1)}{f(q_{i},D)+k_{1}\cdot \left(1-b+b\cdot {\frac {|D|}{\text{avgdl}}}\right)}}}$

${\displaystyle {\text{IDF}}(q_{i})=\ln \left({\frac {N-n(q_{i})+0.5}{n(q_{i})+0.5}}+1\right)}$

${\displaystyle \mathrm {f} (q_{i},D)={\frac {f_{t,d}}{\sum _{t'\in d}{f_{t',d}}}}},$

## CreateTokenScore method variable:

d : ${D}$  
newTotalDoc: ${N}$  
neDocumentNum: ${n(q_{i})}$  
avgdl: ${avgdl}$

k = 1.2  
b = 0.75  
