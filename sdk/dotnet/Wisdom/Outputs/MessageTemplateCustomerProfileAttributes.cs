// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Outputs
{

    /// <summary>
    /// The customer profile attributes that are used with the message template.
    /// </summary>
    [OutputType]
    public sealed class MessageTemplateCustomerProfileAttributes
    {
        /// <summary>
        /// A unique account number that you have given to the customer.
        /// </summary>
        public readonly string? AccountNumber;
        /// <summary>
        /// Any additional information relevant to the customer's profile.
        /// </summary>
        public readonly string? AdditionalInformation;
        /// <summary>
        /// The first line of a customer address.
        /// </summary>
        public readonly string? Address1;
        /// <summary>
        /// The second line of a customer address.
        /// </summary>
        public readonly string? Address2;
        /// <summary>
        /// The third line of a customer address.
        /// </summary>
        public readonly string? Address3;
        /// <summary>
        /// The fourth line of a customer address.
        /// </summary>
        public readonly string? Address4;
        /// <summary>
        /// The first line of a customer’s billing address.
        /// </summary>
        public readonly string? BillingAddress1;
        /// <summary>
        /// The second line of a customer’s billing address.
        /// </summary>
        public readonly string? BillingAddress2;
        /// <summary>
        /// The third line of a customer’s billing address.
        /// </summary>
        public readonly string? BillingAddress3;
        /// <summary>
        /// The fourth line of a customer’s billing address.
        /// </summary>
        public readonly string? BillingAddress4;
        /// <summary>
        /// The city of a customer’s billing address.
        /// </summary>
        public readonly string? BillingCity;
        /// <summary>
        /// The country of a customer’s billing address.
        /// </summary>
        public readonly string? BillingCountry;
        /// <summary>
        /// The county of a customer’s billing address.
        /// </summary>
        public readonly string? BillingCounty;
        /// <summary>
        /// The postal code of a customer’s billing address.
        /// </summary>
        public readonly string? BillingPostalCode;
        /// <summary>
        /// The province of a customer’s billing address.
        /// </summary>
        public readonly string? BillingProvince;
        /// <summary>
        /// The state of a customer’s billing address.
        /// </summary>
        public readonly string? BillingState;
        /// <summary>
        /// The customer's birth date.
        /// </summary>
        public readonly string? BirthDate;
        /// <summary>
        /// The customer's business email address.
        /// </summary>
        public readonly string? BusinessEmailAddress;
        /// <summary>
        /// The name of the customer's business.
        /// </summary>
        public readonly string? BusinessName;
        /// <summary>
        /// The customer's business phone number.
        /// </summary>
        public readonly string? BusinessPhoneNumber;
        /// <summary>
        /// The city in which a customer lives.
        /// </summary>
        public readonly string? City;
        /// <summary>
        /// The country in which a customer lives.
        /// </summary>
        public readonly string? Country;
        /// <summary>
        /// The county in which a customer lives.
        /// </summary>
        public readonly string? County;
        /// <summary>
        /// The custom attributes in customer profile attributes.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Custom;
        /// <summary>
        /// The customer's email address, which has not been specified as a personal or business address.
        /// </summary>
        public readonly string? EmailAddress;
        /// <summary>
        /// The customer's first name.
        /// </summary>
        public readonly string? FirstName;
        /// <summary>
        /// The customer's gender.
        /// </summary>
        public readonly string? Gender;
        /// <summary>
        /// The customer's home phone number.
        /// </summary>
        public readonly string? HomePhoneNumber;
        /// <summary>
        /// The customer's last name.
        /// </summary>
        public readonly string? LastName;
        /// <summary>
        /// The first line of a customer’s mailing address.
        /// </summary>
        public readonly string? MailingAddress1;
        /// <summary>
        /// The second line of a customer’s mailing address.
        /// </summary>
        public readonly string? MailingAddress2;
        /// <summary>
        /// The third line of a customer’s mailing address.
        /// </summary>
        public readonly string? MailingAddress3;
        /// <summary>
        /// The fourth line of a customer’s mailing address.
        /// </summary>
        public readonly string? MailingAddress4;
        /// <summary>
        /// The city of a customer’s mailing address.
        /// </summary>
        public readonly string? MailingCity;
        /// <summary>
        /// The country of a customer’s mailing address.
        /// </summary>
        public readonly string? MailingCountry;
        /// <summary>
        /// The county of a customer’s mailing address.
        /// </summary>
        public readonly string? MailingCounty;
        /// <summary>
        /// The postal code of a customer’s mailing address
        /// </summary>
        public readonly string? MailingPostalCode;
        /// <summary>
        /// The province of a customer’s mailing address.
        /// </summary>
        public readonly string? MailingProvince;
        /// <summary>
        /// The state of a customer’s mailing address.
        /// </summary>
        public readonly string? MailingState;
        /// <summary>
        /// The customer's middle name.
        /// </summary>
        public readonly string? MiddleName;
        /// <summary>
        /// The customer's mobile phone number.
        /// </summary>
        public readonly string? MobilePhoneNumber;
        /// <summary>
        /// The customer's party type.
        /// </summary>
        public readonly string? PartyType;
        /// <summary>
        /// The customer's phone number, which has not been specified as a mobile, home, or business number.
        /// </summary>
        public readonly string? PhoneNumber;
        /// <summary>
        /// The postal code of a customer address.
        /// </summary>
        public readonly string? PostalCode;
        /// <summary>
        /// The ARN of a customer profile.
        /// </summary>
        public readonly string? ProfileArn;
        /// <summary>
        /// The unique identifier of a customer profile.
        /// </summary>
        public readonly string? ProfileId;
        /// <summary>
        /// The province in which a customer lives.
        /// </summary>
        public readonly string? Province;
        /// <summary>
        /// The first line of a customer’s shipping address.
        /// </summary>
        public readonly string? ShippingAddress1;
        /// <summary>
        /// The second line of a customer’s shipping address.
        /// </summary>
        public readonly string? ShippingAddress2;
        /// <summary>
        /// The third line of a customer’s shipping address.
        /// </summary>
        public readonly string? ShippingAddress3;
        /// <summary>
        /// The fourth line of a customer’s shipping address
        /// </summary>
        public readonly string? ShippingAddress4;
        /// <summary>
        /// The city of a customer’s shipping address.
        /// </summary>
        public readonly string? ShippingCity;
        /// <summary>
        /// The country of a customer’s shipping address.
        /// </summary>
        public readonly string? ShippingCountry;
        /// <summary>
        /// The county of a customer’s shipping address.
        /// </summary>
        public readonly string? ShippingCounty;
        /// <summary>
        /// The postal code of a customer’s shipping address.
        /// </summary>
        public readonly string? ShippingPostalCode;
        /// <summary>
        /// The province of a customer’s shipping address.
        /// </summary>
        public readonly string? ShippingProvince;
        /// <summary>
        /// The state of a customer’s shipping address.
        /// </summary>
        public readonly string? ShippingState;
        /// <summary>
        /// The state in which a customer lives.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private MessageTemplateCustomerProfileAttributes(
            string? accountNumber,

            string? additionalInformation,

            string? address1,

            string? address2,

            string? address3,

            string? address4,

            string? billingAddress1,

            string? billingAddress2,

            string? billingAddress3,

            string? billingAddress4,

            string? billingCity,

            string? billingCountry,

            string? billingCounty,

            string? billingPostalCode,

            string? billingProvince,

            string? billingState,

            string? birthDate,

            string? businessEmailAddress,

            string? businessName,

            string? businessPhoneNumber,

            string? city,

            string? country,

            string? county,

            ImmutableDictionary<string, string>? custom,

            string? emailAddress,

            string? firstName,

            string? gender,

            string? homePhoneNumber,

            string? lastName,

            string? mailingAddress1,

            string? mailingAddress2,

            string? mailingAddress3,

            string? mailingAddress4,

            string? mailingCity,

            string? mailingCountry,

            string? mailingCounty,

            string? mailingPostalCode,

            string? mailingProvince,

            string? mailingState,

            string? middleName,

            string? mobilePhoneNumber,

            string? partyType,

            string? phoneNumber,

            string? postalCode,

            string? profileArn,

            string? profileId,

            string? province,

            string? shippingAddress1,

            string? shippingAddress2,

            string? shippingAddress3,

            string? shippingAddress4,

            string? shippingCity,

            string? shippingCountry,

            string? shippingCounty,

            string? shippingPostalCode,

            string? shippingProvince,

            string? shippingState,

            string? state)
        {
            AccountNumber = accountNumber;
            AdditionalInformation = additionalInformation;
            Address1 = address1;
            Address2 = address2;
            Address3 = address3;
            Address4 = address4;
            BillingAddress1 = billingAddress1;
            BillingAddress2 = billingAddress2;
            BillingAddress3 = billingAddress3;
            BillingAddress4 = billingAddress4;
            BillingCity = billingCity;
            BillingCountry = billingCountry;
            BillingCounty = billingCounty;
            BillingPostalCode = billingPostalCode;
            BillingProvince = billingProvince;
            BillingState = billingState;
            BirthDate = birthDate;
            BusinessEmailAddress = businessEmailAddress;
            BusinessName = businessName;
            BusinessPhoneNumber = businessPhoneNumber;
            City = city;
            Country = country;
            County = county;
            Custom = custom;
            EmailAddress = emailAddress;
            FirstName = firstName;
            Gender = gender;
            HomePhoneNumber = homePhoneNumber;
            LastName = lastName;
            MailingAddress1 = mailingAddress1;
            MailingAddress2 = mailingAddress2;
            MailingAddress3 = mailingAddress3;
            MailingAddress4 = mailingAddress4;
            MailingCity = mailingCity;
            MailingCountry = mailingCountry;
            MailingCounty = mailingCounty;
            MailingPostalCode = mailingPostalCode;
            MailingProvince = mailingProvince;
            MailingState = mailingState;
            MiddleName = middleName;
            MobilePhoneNumber = mobilePhoneNumber;
            PartyType = partyType;
            PhoneNumber = phoneNumber;
            PostalCode = postalCode;
            ProfileArn = profileArn;
            ProfileId = profileId;
            Province = province;
            ShippingAddress1 = shippingAddress1;
            ShippingAddress2 = shippingAddress2;
            ShippingAddress3 = shippingAddress3;
            ShippingAddress4 = shippingAddress4;
            ShippingCity = shippingCity;
            ShippingCountry = shippingCountry;
            ShippingCounty = shippingCounty;
            ShippingPostalCode = shippingPostalCode;
            ShippingProvince = shippingProvince;
            ShippingState = shippingState;
            State = state;
        }
    }
}
