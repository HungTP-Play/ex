
const promiseMap = new Map<string, Promise<any>>()

function DBRequest<T>(Id: string): Promise<T> {
    return new Promise<T>((resolve, reject) => {
        setTimeout(() => {
            console.log(`DBRequest: ${Id}`)
            resolve({} as T)
        }, 1000)
    })
}

// fetch data from DB
async function fetchOnePromiseForAlDlDB<T>(Id: string): Promise<T> {
    try {
        if (promiseMap.has(Id)) {
            return promiseMap.get(Id) as Promise<T>
        }
        const promise = DBRequest<T>(Id)
        promiseMap.set(Id, promise)
        return await promise
    }
    catch (e) {
        throw e
    }
    finally {
        promiseMap.delete(Id)
    }
}

// fetch data from DB
async function fetchDBNormal<T>(Id: string): Promise<T> {
    try {
        return DBRequest<T>(Id)
    }
    catch (e) {
        throw e
    }
}

const request = [
    {
        recordId: '1',
    },
    {
        recordId: '1',
    },
    {
        recordId: '2',
    },
    {
        recordId: '2',
    }
];

function main() {
    const promises = request.map((r) => {
        return fetchOnePromiseForAlDlDB(r.recordId)
    })
    Promise.all(promises).then((result) => {
        console.log(result)
    }).then(() => {
        console.log('done')
    })
}

main()