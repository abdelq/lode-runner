# High priority

- [ ] Write a simpler version for TP
- [ ] Broadcast aux spectateurs + gardes Envoyer la carte au complet (Would require out chan in player struct)
- [ ] Broadcast aux Runners Envoyer les déplacements du garde
- [ ] Validate direction LEFT/RIGHT for digging actions
- [ ] Touching special ladder means going to next level + no hiding bc of money, only allowing to go to next lvl
- [ ] Action type should be a const (MOVE DIG) used in update action
- [ ] Leave and join back same room (room not deleted)
- [ ] Broadcast aux Runners (seulement la nouvelle position du runner)
- [ ] Ne pas marcher sur la rope (lvl 007)
- [ ] Arriver sur un ladder depuis une rope fait tomber moment, éventuellement, ça se chie dessus
- [ ] runner.go, guard.go
- [ ] La sortie existe dès le début et ramasser tous le gold la fait disparaître

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
- [ ] Flags for timeouts
- [ ] Ne pas pouvoir monter escalier caché quand il apparaît
- [ ] Tomber à travers une rope au lieu de la pogner
- [ ] Arriver sur un ladder depuis une rope fait tomber
