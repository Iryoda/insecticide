schema "public"  {
enum "action"  {
    schema = schema.public
    values = ["click", "doubleclick", "text", "goto"]
}

enum "assertion"  {
    schema = schema.public
    values = ["text"]
}

table "step" {
    schema = schema.public
    column "id" {
        type = uuid
        default = sql("gen_random_uuid()")
    }
    column "name" {
        type = varchar(255)
    }
    column "isUnique" {
        type = boolean
        default = false 
    }
    column "selector" {
        type = varchar(255)
    }
    column "componentId" {
        type = uuid
        null= true 
    }
    column "nextStepId" {
        type = uuid
        null = true
    }
    column "action" {
        type = enum.action 
        null = true
    }
    column "assertion" {
        type = enum.assertion 
        null = true
    }
    column "assertionTarget" {
        type = varchar(255) 
        null = true
    }
    column "createdAt" {
        type = timestamptz 
        default = sql("now()")
    }
    column "updatedAt" {
        type = timestamptz  
        null = true
    } 
    primary_key {
        columns = [column.id]
    }
}
