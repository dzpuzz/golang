package tree

import (
	"math"
)

//二叉树节点
type BinTreeNode struct {
	data   interface{}  //数据域
	parent *BinTreeNode //父节点
	lChild *BinTreeNode //左孩子
	rChild *BinTreeNode //右孩子
	height int          //以该结点为根的子树的高度
	size   int          //该结点子孙数(包括结点本身)
}

//新建节点
func NewBinTreeNode(e interface{}) *BinTreeNode {
	return &BinTreeNode{data: e, size: 1}
}

//获得节点数据
func (this *BinTreeNode) GetData() interface{} {
	if this == nil {
		return nil
	}
	return this.data
}

//设置节点数据
func (this *BinTreeNode) SetData(e interface{}) {
	this.data = e
}

//判断是否有父亲
func (this *BinTreeNode) HasParent() bool {
	return this.parent != nil
}

//获得父亲节点
func (this *BinTreeNode) GetParent() *BinTreeNode {
	if !this.HasParent() {
		return nil
	}
	return this.parent
}

//设置父亲节点
func (this *BinTreeNode) SetParent(p *BinTreeNode) {
	this.parent = p
}

//断开与父亲的关系
func (this *BinTreeNode) CutOffParent() {
	if !this.HasParent() {
		return
	}
	if this.IsLChild() {
		this.parent.lChild = nil //断开该节点与父节点的连接
	} else {
		this.parent.rChild = nil //断开该节点与父节点的连接
	}

	this.parent = nil       //断开该节点与父节点的连接
	this.parent.SetHeight() //更新父结点及其祖先高度
	this.parent.SetSize()   //更新父结点及其祖先规模
}

//判断是否为此节点左孩子
func (this *BinTreeNode) IsLChild() bool {
	return this.HasParent() && this == this.parent.lChild
}

//判断是否有左孩子
func (this *BinTreeNode) HasLChild() bool {
	return this.lChild != nil
}

//判断是否有右孩子
func (this *BinTreeNode) HasRChild() bool {
	return this.rChild != nil
}

//返回左节点
func (this *BinTreeNode) getLChild() *BinTreeNode {
	if !this.HasLChild() {
		return nil
	}
	return this.lChild
}

//返回右节点
func (this *BinTreeNode) getRChild() *BinTreeNode {
	if !this.HasRChild() {
		return nil
	}
	return this.rChild
}

//判断是否为此节点右孩子
func (this *BinTreeNode) IsRChild() bool {
	return this.HasParent() && this == this.parent.rChild
}

//判断是否为叶子结点
func (this *BinTreeNode) IsLeaf() bool {
	return !this.HasLChild() && !this.HasRChild()
}

//获得以此为根的高度
func (this *BinTreeNode) GetHeight() int {
	return this.height
}

//设置当前结点的左孩子,返回原左孩子
func (this *BinTreeNode) SetLChild(lc *BinTreeNode) *BinTreeNode {
	oldLC := this.lChild
	if this.HasLChild() {
		this.lChild.CutOffParent() //断开当前左孩子与结点的关系
	}
	if lc != nil {
		lc.CutOffParent() //断开lc与其父结点的关系
		this.lChild = lc  //确定父子关系
		lc.setParent(this)
		this.SetHeight() //更新当前结点及其祖先高度
		this.SetSize()   //更新当前结点及其祖先规模
	}
	return oldLC
}

//设置当前节点的右孩子,返回原右孩子
func (this *BinTreeNode) SetRChild(rc *BinTreeNode) *BinTreeNode {
	oldRC := this.rChild
	if this.HasRChild() {
		this.rChild.CutOffParent()
	}
	if rc != nil {
		rc.CutOffParent()
		this.rChild = rc
		lc.setParent(this)
		this.SetHeight()
		this.SetSize()
	}
	return oldRC
}
