import TreeNode from './TreeNode'
import chalk from 'chalk'

describe('TreeNode', () => {
  it('root node has no child nodes', () => {
    const root = new TreeNode('root')

    expect(root.nodes).toHaveLength(0)
  })

  it('root nodes has 3 child nodes', () => {
    const root = new TreeNode('root')

    root.node = new TreeNode('hey')
    root.node = new TreeNode('whiskey tango foxtrot')
    root.node = new TreeNode('foxy mama')

    expect(root.nodes).toHaveLength(3)
  })

  it('#isLeaf returns false', () => {
    const root = new TreeNode('root')
    root.node = new TreeNode('hey')

    expect(root.isLeaf()).toBe(false)
  })

  it('#isLeaf returns true', () => {
    const root = new TreeNode('root')

    expect(root.isLeaf()).toBe(true)
  })
})
