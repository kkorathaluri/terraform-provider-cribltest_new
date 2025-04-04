// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type GetInvoicesResponseV5 struct {
	Items []InvoiceV5 `json:"items"`
	Count float64     `json:"count"`
}

func (o *GetInvoicesResponseV5) GetItems() []InvoiceV5 {
	if o == nil {
		return []InvoiceV5{}
	}
	return o.Items
}

func (o *GetInvoicesResponseV5) GetCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Count
}
