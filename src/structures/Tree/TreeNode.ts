import cuid from 'cuid'

export default class TreeNode {
  private _id: string = cuid()
  private _data?: string
  private _nodes: TreeNode[] = []

  constructor(data?: string) {
    this._data = data
  }

  public get nodes(): TreeNode[] {
    return this._nodes
  }

  public set node(node: TreeNode) {
    this._nodes.push(node)
  }

  public get data() {
    return this._data
  }

  public isLeaf(): boolean {
    return this._nodes.length === 0
  }
}
