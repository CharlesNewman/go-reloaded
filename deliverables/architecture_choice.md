### Architecture choice: Pipeline

This program uses a pipeline architecture because the text is processed through a fixed sequence of stages.

First, the program reads the input file. Then it tokenizes the text, applies transformation commands, fixes articles, rebuilds the final text, and writes the result to the output file.

This structure fits the project well because each stage has one clear responsibility and passes its result to the next stage.

The program is not mainly an FSM, because the overall design is not based on switching between many states. Some small parts, like quote handling, use state-like logic, but the full architecture is still a pipeline.