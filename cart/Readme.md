## Cart Package ##

### Configurations###

```
  
  cart:
    # To register the in memory cart service adapter (e.g. for development mode)
    useInMemoryCartServiceAdapters: true
    
```

### Cart API ###

Adding simple products:

http://localhost:3210/en/api/cart/add?marketplaceCode=fake_simple

With qty:
http://localhost:3210/en/api/cart/add?marketplaceCode=fake_simple&qty=10

Adding configurables:
http://localhost:3210/en/api/cart/add?marketplaceCode=fake_configurable&variantMarketplaceCode=shirt-white-s