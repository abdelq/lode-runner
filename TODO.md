# High priority

- [x] Use a system based on ticks
- [x] game.Stop() n'arrête pas le jeu
- [x] Manage death
    - [x] Manage collisions
    - [x] Quand plus de vies
    - [x] Fin de partie quand le joueur tombe dans un trou au premier étage
- [ ] Broadcast aux spectateurs + gardes
      Envoyer la carte au complet
- [ ] Broadcast aux Runners
      Envoyer les déplacements du garde
- [ ] Ne pas marcher sur la rope (lvl 007)
- [ ] Arriver sur un ladder depuis une rope fait tomber
- [ ] Bug weird : laisser runner le serveur + client random pour un
  moment, éventuellement, ça se chie dessus
- [ ] Validate direction LEFT/RIGHT for digging actions
- [ ] Disable guard from game (game.filled()) + Update quit message + Comment ./run.sh for guard

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
- [ ] Let's Encrypt certs
- [ ] Clean up JS client
