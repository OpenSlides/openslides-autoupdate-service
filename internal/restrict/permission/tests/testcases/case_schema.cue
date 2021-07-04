{#Case}

#Case: {
    name?: string

    // TODO: special forms allowed.
    db?:  [string]: _

    // TODO: check values.
    fqfields?: [...string]
    fqids?: [...string]

    user_id: >=0 | *1337

    meeting_id: >0 | *1337

    permission?: string

    // TODO: has to be fqid or fqfields.
    can_see?: [...string]
    can_not_see?: [...string]

    cases?: [...#Case]
}
