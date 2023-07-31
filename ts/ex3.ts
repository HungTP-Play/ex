type Item = {
    weight: number;
    value: number;
}

function ex3KnapsackRecursive(items: Item[], capacity: number): number {
    const knapsack = (items: Item[], capacity: number, currentWeight: number, currentValue: number, index: number): number => {
        if (currentWeight > capacity) {
            return 0;
        }

        if (index === items.length) {
            return currentValue;
        }

        // With current item
        selectedItems[index] = true;
        const withCurrentItem: number = knapsack(items, capacity, currentWeight + items[index].weight, currentValue + items[index].value, index + 1);

        // Without current item
        selectedItems[index] = false;
        const withoutCurrentItem: number = knapsack(items, capacity, currentWeight, currentValue, index + 1);

        return Math.max(withCurrentItem, withoutCurrentItem);
    }

    const selectedItems = new Array(items.length).fill(false);
    return knapsack(items, capacity, 0, 0, 0);
}

function ex3KnapsackDynamic(items: Item[], capacity: number): number {
    const itemCount = items.length;
    const knapsack: number[][] = new Array(itemCount + 1).fill(0).map(() => new Array(capacity + 1).fill(0));

    for (let i = 1 ; i <= itemCount ; i++) {
        for (let j=1 ; j <= capacity ; j++) {
            const itemIndex = i - 1;
            const {weight, value} = items[itemIndex];
            if (weight > j) {
                // Ignore item
                knapsack[i][j] = knapsack[i-1][j];
            } else {
                // Take max between with item and without item
                const maxWithoutItem = knapsack[i-1][j];
                const maxWithItem = value + knapsack[i-1][j-weight];
                knapsack[i][j] = Math.max(maxWithoutItem, maxWithItem);
            }
        }
    }

    return knapsack[itemCount][capacity];
}

function ex3main(){
    const items: Item[] = [
        {weight: 2, value: 6},
        {weight: 2, value: 10},
        {weight: 3, value: 12},
    ];
    const capacity = 5;

    console.log(ex3KnapsackRecursive(items, capacity));
    console.log(ex3KnapsackDynamic(items, capacity));
}

ex3main();