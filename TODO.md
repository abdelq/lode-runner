# High priority

- [ ] Flags for timeouts
- [ ] Ne pas pouvoir monter escalier caché quand il apparaît
- [ ] Tomber à travers une rope au lieu de la pogner
- [x] Use a system based on ticks
- [x] game.Stop() n'arrête pas le jeu
- [x] Manage death
    - [x] Manage collisions
    - [x] Quand plus de vies
    - [x] Fin de partie quand le joueur tombe dans un trou au premier étage
- [ ] Leave and join back same room (room not deleted)
- [x] Broadcast aux spectateurs + gardes
      Envoyer la carte au complet
- [ ] Ne pas marcher sur la rope (lvl 007)
- [ ] Arriver sur un ladder depuis une rope fait tomber
- [x] Bug weird : laisser runner le serveur + client random pour un
  moment, éventuellement, ça se chie dessus
- [ ] Validate direction LEFT/RIGHT for digging actions
- [x] Disable guard from game (just comment section from and stop the join game.filled()) + Update quit message + Comment ./run.sh for guard
- [x] TCP over browser...
- [ ] Write a simpler version for TP
- [x] Broadcast aux spectateurs + gardes Envoyer la carte au complet
- [ ] Broadcast aux Runners (seulement la nouvelle position du runner) (Would require out chan in player struct)
- [ ] Touching special ladder means going to next level + no hiding bc of money, only allowing to go to next lvl
- [ ] Action type should be a const (MOVE DIG) used in update action
- [ ] runner.go, guard.go
- [ ] La sortie existe dès le début et ramasser tout le gold la fait disparaître
- [ ] Donner plus de tics à un trou avant de se refermer (une seule == impossible de tomber dedans). 3 ou 4 tics ça serait legit
- [ ] Impossible d'aller sur la dernière colonne dans le level de TP 001.lvl

# Mid priority

- [ ] Fix unit tests + finish them
- [ ] Turns out que même si y'a pas de blocs, le bas compte comme du sol (lvl 014)
- [ ] X tiles (false blocks)
- [ ] Respawn of guards

# Low priority

- [ ] Essentiellement tous les niveaux passé lvl 7 (inclus)
- [ ] End of game (lvl 150)
- [ ] Allow multiple rooms for a single client
- [ ] Proper makefiles + .gitignore
- [ ] Clean up JS client
- [ ] TLS as option (when using 443)
- [ ] Réactiver le guarde
    - [ ] Envoyer les déplacements du garde au runner
