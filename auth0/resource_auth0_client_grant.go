package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yieldr/go-auth0/management"
)

func newClientGrant() *schema.Resource {
	return &schema.Resource{

		Create: createClientGrant,
		Read:   readClientGrant,
		Update: updateClientGrant,
		Delete: deleteClientGrant,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"audience": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scope": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func createClientGrant(d *schema.ResourceData, m interface{}) error {
	g := buildClientGrant(d)
	api := m.(*management.Management)
	if err := api.ClientGrant.Create(g); err != nil {
		return err
	}
	d.SetId(g.ID)
	return readClientGrant(d, m)
}

func readClientGrant(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	g, err := api.ClientGrant.Read(d.Id())
	if err != nil {
		return err
	}
	d.SetId(g.ID)
	d.Set("client_id", g.ClientID)
	d.Set("audience", g.Audience)
	d.Set("scope", g.Scope)
	return nil
}

func updateClientGrant(d *schema.ResourceData, m interface{}) error {
	g := buildClientGrant(d)
	g.Audience = ""
	g.ClientID = ""
	api := m.(*management.Management)
	err := api.ClientGrant.Update(d.Id(), g)
	if err != nil {
		return err
	}
	return readClientGrant(d, m)
}

func deleteClientGrant(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.ClientGrant.Delete(d.Id())
}

func buildClientGrant(d *schema.ResourceData) *management.ClientGrant {
	g := &management.ClientGrant{
		ClientID: d.Get("client_id").(string),
		Audience: d.Get("audience").(string),
		Scope:    d.Get("scope").([]interface{}),
	}

	return g
}
