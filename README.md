# Go Reloaded

Outil en ligne de commande, écrit en Go, qui corrige automatiquement un fichier texte.
Projet réalisé à **Zone 01 Rouen** (titre Concepteur Développeur d'Applications).

## Ce que fait le programme

Il lit un fichier, applique les règles suivantes, puis écrit le résultat dans un autre fichier :

| Dans le texte | Résultat |
| --- | --- |
| `1E (hex)` | le nombre hexadécimal est converti en décimal : `30` |
| `10 (bin)` | le nombre binaire est converti en décimal : `2` |
| `mot (up)` / `(low)` / `(cap)` | le mot passe en majuscules, minuscules ou avec une majuscule initiale |
| `(up, 2)` | la règle s'applique aux 2 mots précédents |
| `hello ,world !` | la ponctuation est recollée : `hello, world!` |
| `' mot '` | les apostrophes sont recollées autour du mot |
| `a apple` | `a` devient `an` devant une voyelle ou un h |

Les transformations utilisent des expressions régulières (`regexp`).

## Utilisation

```bash
go run . sample.txt result.txt
```

## Auteur

Abdou Cherif Hono Digo – [portfolio](https://acherif88.github.io)
