{#Case}

#Case: {
    name?: string

    db?: [#Collection]: [#ID]: [#Field]: _
    db?: [#FQID]: [#Field]: _
    db?: [#FQField]: _

    fqfields?: [...#FQField]
    fqids?: [...#FQID]

    user_id: int & >=0 | *1337
    meeting_id: int & >0 | *1337

    permission?: string

    can_see?: [...#FQField | #FQID]
    can_not_see?: [...#FQField | #FQID]

    cases?: [...#Case]
}

#Collection: =~ #"^([a-z]+|[a-z][a-z_]*[a-z])$"#
#ID: =~ #"^[1-9][0-9]*$"#
#Field: =~ #"^[a-z][a-z0-9_]*\$?[a-z0-9_]*$"#
#FQID: =~ #"^([a-z]+|[a-z][a-z_]*[a-z])/[1-9][0-9]*$"#
#FQField: =~ #"^([a-z]+|[a-z][a-z_]*[a-z])/[1-9][0-9]*/[a-z][a-z0-9_]*\$?[a-z0-9_]*$"#
