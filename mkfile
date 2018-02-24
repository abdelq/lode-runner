all: directives.pdf bundle.zip

bundle.zip:V: directives.pdf
    zip -r bundle.zip directives.pdf clients/NodeRunnerKeyboard/ clients/NodeRunnerAI/ clients/javascript/

&.pdf: &.md
    pandoc $stem.md \
    -Vgeometry:margin=1.5in \
    --variable linkcolor=blue \
    -o $stem.pdf
