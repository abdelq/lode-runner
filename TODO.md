# High priority

- [ ] client/ game/
- [ ] Write simpler version for TP
- [ ] Fix unit tests + finish them
- [ ] Leave and join back same room (room not deleted)
- [ ] Broadcast aux spectateurs + gardes
      Envoyer la carte au complet
- [ ] Broadcast aux Runners
      Envoyer les déplacements du garde
- [ ] Ne pas marcher sur la rope (lvl 007)
- [ ] Arriver sur un ladder depuis une rope fait tomber
- [ ] Validate direction LEFT/RIGHT for digging actions
- [ ] Disable guard from game (just comment section from and stop the join game.filled()) + Update quit message + Comment ./run.sh for guard
- [ ] TCP over browser...
- [ ] TLS as option (when using 443)
- [ ] Touching special ladder means going to next level + no hiding bc of money, only allowing to go to next lvl
- [ ] Action type should be a const (MOVE DIG) used in update action

# Mid priority

- [ ] Turns out que même si y'a pas de blocs, le bas compte comme du sol (lvl 014)
- [ ] X tiles (false blocks)
- [ ] Respawn of guards

# Low priority

- [ ] Essentiellement tous les niveaux passé lvl 7 (inclus)
- [ ] End of game (lvl 150)
- [ ] Allow multiple rooms for a single client
- [ ] Flags for timeouts
- [ ] Proper makefiles + .gitignore
- [ ] Clean up JS client
- [ ] Try connect 443 then 1337 (client)
