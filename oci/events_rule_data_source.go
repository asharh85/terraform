// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_events "github.com/oracle/oci-go-sdk/events"
)

func EventsRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["rule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(EventsRuleResource(), fieldMap, readSingularEventsRule)
}

func readSingularEventsRule(d *schema.ResourceData, m interface{}) error {
	sync := &EventsRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).eventsClient

	return ReadResource(sync)
}

type EventsRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_events.EventsClient
	Res    *oci_events.GetRuleResponse
}

func (s *EventsRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EventsRuleDataSourceCrud) Get() error {
	request := oci_events.GetRuleRequest{}

	if ruleId, ok := s.D.GetOkExists("rule_id"); ok {
		tmp := ruleId.(string)
		request.RuleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "events")

	response, err := s.Client.GetRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *EventsRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Actions != nil {
		s.D.Set("actions", []interface{}{ActionListToMap(s.Res.Actions, true)})
	} else {
		s.D.Set("actions", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Condition != nil {
		s.D.Set("condition", *s.Res.Condition)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.LifecycleMessage != nil {
		s.D.Set("lifecycle_message", *s.Res.LifecycleMessage)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}