# High priority

- [x] Use a system based on ticks
- [x] game.Stop() n'arrête pas le jeu
- [x] Manage death
    - [x] Manage collisions
    - [x] Quand plus de vies
    - [x] Fin de partie quand le joueur tombe dans un trou au premier étage
- [ ] Leave and join back same room (room not deleted)
- [x] Broadcast aux spectateurs + gardes
      Envoyer la carte au complet
- [ ] Broadcast aux Runners (seulement la nouvelle position du runner)
- [ ] Ne pas marcher sur la rope (lvl 007)
- [ ] Arriver sur un ladder depuis une rope fait tomber
- [x] Bug weird : laisser runner le serveur + client random pour un
  moment, éventuellement, ça se chie dessus
- [ ] Validate direction LEFT/RIGHT for digging actions
- [x] Disable guard from game (just comment section from and stop the join game.filled()) + Update quit message + Comment ./run.sh for guard
- [ ] runner.go, guard.go
- [x] TCP over browser...
- [ ] La sortie existe dès le début et ramasser tous le gold la fait disparaître

# Mid priority

- [ ] Turns out que même si y'a pas de blocs, le bas compte comme du sol (lvl 014)
- [ ] X tiles (false blocks)
- [x] Pouvoir descendre des échelles
- [x] Besoin de taper deux fois sur une direction lorsqu'on est sur une rope
- [x] Ne pas avoir accès à l'échelle de sortie avant la fin
- [x] Lvl 01: impossible de ramasser le tout premier $
- [x] Death on block rebuild
- [x] Manage going to the next level
- [ ] Respawn of guards

# Low priority

- [ ] Essentiellement tous les niveaux passé lvl 7 (inclus)
- [ ] End of game (lvl 150)
- [ ] Allow multiple rooms for a single client
- [ ] Flags for timeouts
- [x] Permettre de tomber d'une ROPE
- [ ] Proper makefiles + .gitignore
- [ ] Clean up JS client
- [ ] TLS as option (when using 443)
- [ ] Réactiver le guarde
    - [ ] Envoyer les déplacements du garde au runner
