import protoLoader from '@grpc/proto-loader'
import grpc from '@grpc/grpc-js'

const ProtoPath = './proto/catalogue.proto'
const packageDefinition = protoLoader.loadSync(ProtoPath)
const catalogue_proto = grpc.loadPackageDefinition(packageDefinition).book

const clientUri = "localhost:50051"
console.log(clientUri)

const client = new catalogue_proto.Catalogue(
    clientUri, grpc.credentials.createInsecure()
)

export class CatalogueDataSource {
    constructor() {
        this.client = client
    }
    async getBook(id) {
        return new Promise((resolve, reject) => {
            this.client.getBook({id: id}, (error, response) => {
                if (error) {
                    return reject(err)
                } else {
                    return resolve(response)
                }
            })
        })
    }

    async listBooks() {
        return new Promise((resolve, reject) => {
            this.client.listBooks({}, (error, response) => {
                if (error) {
                    return reject(err)
                } else {
                    return resolve(response)
                }
            })
        })
    }
}
