# Notes

- Explainers
    - The 'Strategy' type has been modelled such so that we can have strategies other than 'Holding at number x'. Eg. Hold until two fives occur
    - We can probably abstract out the loop from the main function into another function as well
    - I thought of stories as incremental improvements so previous versions are in my commit history

- Improvements
    - Make tests more readable with names, right now they just read as gibberish
    - Maybe granularize functions further eg. 'GetDecision' for player
    - More explicit errors
    - Use cobra for cli