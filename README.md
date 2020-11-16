# asm6csv

Der Apple School Manager (ASM) importiert Accounts aus CSV-Dateien. Die Eigenschaften und Beziehungen werden in 6 verschiedene Tabellen abgebildet. Apple [dokumentiert den Aufbau](https://support.apple.com/de-de/HT207029) ziemlich ausführlich. Die Bezüge zwischen den Spalten verschiedener Tabellen müssen dabei exakt eingehalten werden. Das macht das händische Anlegen dieser 6 Tabellen mühsam.

Da wir die interne Struktur der Schule nicht im ASM abbilden wollen, aber für die Nutzung von verwalteten Apple Ids die Accounts brauchen, beschränken wir uns auf eine gemischte Tabelle, deren Inhalt dieses Tool auf die benötigten 6 Tabellen verteilt.

## Kommandozeilenargumente

  ```shell
  $ ./asm6csv -h
Usage: asm6csv [options] filename.csv
Options:
  -t    Generate csv template (shorthand).
  -template
        Generate csv template
```

## Beispiel für eine ausgefüllte Tabelle
| student_id | first_name | last_name | class_id | course_id | teacher_id | password_policy | location_name |
|---|---|---|---|---|---|---|---|
| blau1 | blau | 1 | Blau | Blau | Lempel |  | Schule |
| blau2 | blau | 2 | Blau | Blau | Lempel |  | Schule |
| blau3 | blau | 3 | Blau | Blau | Lempel |  | Schule |
| blau4 | blau | 4 | Blau | Blau | Lempel |  | Schule |
| rot1 | rot | 1 | Rot | Rot | Lempel |  | Schule |