# Receipt Rule-Engine Challenge

Implement the `NewProcessor()` function in the `processor` package of this repository.

Implement the two rules defined under the Rules section below, so they can be passed into the `addRule` function.

> NOTE
> 
> The code in the `processor` and `model` packages will be used to test your submission. Editing these files other
> than adding your initialization code to this function risks breaking our test suite that will evaluate your submission.
> If this happens, your application will be rejected before it is reviewed by one of our engineers.

## Context

This system is part of a hypothetical receipt processing system. When a user submits a receipt it is first normalized.
This normalization ensures the following statements are true:

* The following Receipt fields are populated with non-blank values
  * Receipt ID
  * StoreName
  * PurchaseTime
  * Items Array
* All items have prices greater than $0.00

The normalization process also performs a product lookup against our internal catalog. The ID value for each item is
an internal Fetch ID that clearly identifies what the item purchased is. However, this product matching isn't perfect.
If a suitable ID cannot be found, this ID will be blank.

Because this normalization is happening, receipts passed into the system can be trusted to follow these guidelines.

However, rules are provided by internal Fetch users. Inputs to the `AddRule` function should consider the rule 
definitions to be user input and handled accordingly.

# Rules

Rules represent features requested by our Product team. When the processor is initialized, it will be provided a set
of rules. When receipts are later processed, they are evaluated against these rules to determine how many points should
be awarded.

All rules are processed against each receipt.

## Store Name

`StoreName` awards a flat amount of points if the user shops at the specified store.

The rule provides a value. If the store name contains the value as a case-insensitive sub-string, it is considered a match.

For example, the value `buck` would match both `Starbucks` and `Bucky's Tavern`

The rule either matches, or it doesn't. It cannot match more than once.

Example Rule Definition:

```json
{"value":"Target","points":100}
```

### Rule Attributes

#### Value

Value must be defined.

#### Points

Points must be a positive integer.

## Item Match

`ItemMatch` awards points as a percentage of the price of the matched item. These items are matched regardless of where
they are purchased. When defining the rule, a list of all matching ID's will be provided along with a rate multiplier
such that `1.0` represents 100%. 

For purposes of this calculation, $1.00 is assumed to be equivalent to 1000 points.

For example if we are awarding 10% points on item `1234`, then we would receive this rule definition:

```json
{"ids":["1234"],"rate":0.1}
```

Then if we get a receipt where item `1234` is purchased for $12.25, then we would award 10% of 12250 points or 1225 points.

If this calculation results in fractional points, the awarded points should be rounded down.
